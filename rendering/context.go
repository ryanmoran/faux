package rendering

import (
	"fmt"

	"github.com/ryanmoran/faux/parsing"
)

type Context struct {
	imports Imports
}

func NewContext() *Context {
	return &Context{
		imports: Imports{
			Package{Path: "sync"},
		},
	}
}

func (c *Context) Build(iface parsing.Interface) File {
	fake := c.BuildFakeType(iface)

	var funcs []Func
	for _, signature := range iface.Signatures {
		funcs = append(funcs, c.BuildFunc(fake, signature))
	}

	return NewFile("fakes", c.imports, []NamedType{fake}, funcs)
}

func (c *Context) BuildFakeType(iface parsing.Interface) NamedType {
	var calls []Field

	for _, signature := range iface.Signatures {
		calls = append(calls, c.BuildCallStruct(signature))
	}

	return NewNamedType(TitleString(iface.Name), NewStruct(calls))
}

func (c *Context) BuildCallStruct(signature parsing.Signature) Field {
	methodCallName := fmt.Sprintf("%sCall", signature.Name)
	var fields []Field
	fields = append(fields, c.BuildMutex())
	fields = append(fields, c.BuildCallCount())

	if len(signature.Params) > 0 {
		fields = append(fields, c.BuildReceives(signature.Params))
	}

	if len(signature.Results) > 0 {
		fields = append(fields, c.BuildReturns(signature.Results))
	}

	fields = append(fields, c.BuildStub(signature))

	return NewField(methodCallName, NewStruct(fields))
}

func (c *Context) BuildMutex() Field {
	return NewField("", NewNamedType("sync.Mutex", NewStruct(nil)))
}

func (c *Context) BuildCallCount() Field {
	return NewField("CallCount", NewBasicType(BasicInt))
}

func (c *Context) BuildReceives(args []parsing.Argument) Field {
	var fields []Field
	for i, arg := range args {
		if arg.Package != nil {
			c.imports.Add(arg.Package)
		}

		name := arg.Name
		if name == "" {
			name = FieldTypeName(args, i)
		}

		field := NewField(name, NewType(arg.Type))
		fields = append(fields, field)
	}

	return NewField("Receives", NewStruct(fields))
}

func (c *Context) BuildReturns(args []parsing.Argument) Field {
	var fields []Field
	for i, arg := range args {
		if arg.Package != nil {
			c.imports.Add(arg.Package)
		}

		name := arg.Name
		if name == "" {
			name = FieldTypeName(args, i)
		}

		field := NewField(name, NewType(arg.Type))
		fields = append(fields, field)
	}

	return NewField("Returns", NewStruct(fields))
}

func (c *Context) BuildStub(signature parsing.Signature) Field {
	params := c.BuildParams(signature.Params, false)
	results := c.BuildResults(signature.Results)
	stub := NewFunc("", Receiver{}, params, results, nil)
	return NewField("Stub", stub)
}

func (c *Context) BuildFunc(fake NamedType, signature parsing.Signature) Func {
	receiver := NewReceiver("f", NewPointer(fake))
	params := c.BuildParams(signature.Params, true)
	results := c.BuildResults(signature.Results)
	body := c.BuildBody(receiver, signature)

	return NewFunc(signature.Name, receiver, params, results, body)
}

func (c *Context) BuildParams(args []parsing.Argument, named bool) []Param {
	var params []Param
	for i, arg := range args {
		var name string
		if named {
			name = ParamName(i)
		}

		params = append(params, NewParam(name, NewType(arg.Type), arg.Variadic))
	}

	return params
}

func (c *Context) BuildResults(args []parsing.Argument) []Result {
	var results []Result
	for _, arg := range args {
		results = append(results, NewResult(NewType(arg.Type)))
	}

	return results
}

func (c *Context) BuildBody(receiver Receiver, signature parsing.Signature) []Statement {
	statements := []Statement{
		c.BuildMutexLockStatement(receiver, signature.Name),
		c.BuildMutexUnlockStatement(receiver, signature.Name),
		c.BuildIncrementStatement(receiver, signature.Name),
	}

	for i, _ := range signature.Params {
		statements = append(statements, c.BuildAssignStatement(receiver, signature.Name, i, signature.Params))
	}

	statements = append(statements, c.BuildStubIfStatement(receiver, signature))

	if len(signature.Results) > 0 {
		statements = append(statements, c.BuildReturnStatement(receiver, signature))
	}

	return statements
}

func (c *Context) BuildMutexLockStatement(receiver Receiver, name string) CallStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", name))
	selector := NewSelector(receiver, callField, NewIdent("Lock"))

	return NewCallStatement(NewCall(selector))
}

func (c *Context) BuildMutexUnlockStatement(receiver Receiver, name string) DeferStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", name))
	selector := NewSelector(receiver, callField, NewIdent("Unlock"))

	return NewDeferStatement(NewCall(selector))
}

func (c *Context) BuildIncrementStatement(receiver Receiver, name string) IncrementStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", name))
	countField := callField.Type.(Struct).FieldWithName("CallCount")
	selector := NewSelector(receiver, callField, countField)

	return NewIncrementStatement(selector)
}

func (c *Context) BuildAssignStatement(receiver Receiver, name string, index int, args []parsing.Argument) AssignStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", name))
	receivesField := callField.Type.(Struct).FieldWithName("Receives")

	arg := args[index]
	argName := arg.Name
	if argName == "" {
		argName = FieldTypeName(args, index)
	}

	paramField := receivesField.Type.(Struct).FieldWithName(argName)
	selector := NewSelector(receiver, callField, receivesField, paramField)
	paramName := ParamName(index)
	param := NewParam(paramName, NewType(arg.Type), arg.Variadic)

	return NewAssignStatement(selector, param)
}

func (c *Context) BuildStubIfStatement(receiver Receiver, signature parsing.Signature) IfStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", signature.Name))
	stubField := callField.Type.(Struct).FieldWithName("Stub")
	selector := NewSelector(receiver, callField, stubField)

	params := c.BuildParams(signature.Params, true)

	condition := NewEquality(false, selector, NewNil())
	var body []Statement

	if len(signature.Results) > 0 {
		body = append(body, NewReturnStatement(NewCall(selector, params...)))
	} else {
		body = append(body, NewCallStatement(NewCall(selector, params...)))
	}

	return NewIfStatement(condition, body)
}

func (c *Context) BuildReturnStatement(receiver Receiver, signature parsing.Signature) ReturnStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", signature.Name))
	returnsField := callField.Type.(Struct).FieldWithName("Returns")

	var results []Expression
	for i, arg := range signature.Results {
		argName := arg.Name
		if argName == "" {
			argName = FieldTypeName(signature.Results, i)
		}

		resultField := returnsField.Type.(Struct).FieldWithName(argName)
		selector := NewSelector(receiver, callField, returnsField, resultField)
		results = append(results, selector)
	}

	return NewReturnStatement(results...)
}
