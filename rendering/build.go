package rendering

import (
	"fmt"

	"github.com/ryanmoran/faux/parsing"
)

func Build(iface parsing.Interface) File {
	fake := BuildFakeType(iface)

	var funcs []Func
	for _, signature := range iface.Signatures {
		funcs = append(funcs, BuildFunc(fake, signature))
	}

	return NewFile("fakes", []NamedType{fake}, funcs)
}

func BuildFakeType(iface parsing.Interface) NamedType {
	var calls []Field

	for _, signature := range iface.Signatures {
		calls = append(calls, BuildCallStruct(signature))
	}

	return NewNamedType(TitleString(iface.Name), NewStruct(calls))
}

func BuildCallStruct(signature parsing.Signature) Field {
	methodCallName := fmt.Sprintf("%sCall", signature.Name)
	var fields []Field
	fields = append(fields, BuildCallCount())

	if len(signature.Params) > 0 {
		fields = append(fields, BuildReceives(signature.Params))
	}

	if len(signature.Results) > 0 {
		fields = append(fields, BuildReturns(signature.Results))
	}

	return NewField(methodCallName, NewStruct(fields))
}

func BuildCallCount() Field {
	return NewField("CallCount", NewBasicType(BasicInt))
}

func BuildReceives(args []parsing.Argument) Field {
	var fields []Field
	for i, arg := range args {
		name := arg.Name
		if name == "" {
			name = FieldTypeName(args, i)
		}

		field := NewField(name, NewType(arg.Type))
		fields = append(fields, field)
	}

	return NewField("Receives", NewStruct(fields))
}

func BuildReturns(args []parsing.Argument) Field {
	var fields []Field
	for i, arg := range args {
		name := arg.Name
		if name == "" {
			name = FieldTypeName(args, i)
		}

		field := NewField(name, NewType(arg.Type))
		fields = append(fields, field)
	}

	return NewField("Returns", NewStruct(fields))
}

func BuildFunc(fake NamedType, signature parsing.Signature) Func {
	receiver := NewReceiver("f", NewPointer(fake))
	params := BuildParams(signature.Params)
	results := BuildResults(signature.Results)
	body := BuildBody(receiver, signature)

	return NewFunc(signature.Name, receiver, params, results, body)
}

func BuildParams(args []parsing.Argument) []Param {
	var params []Param
	for i, arg := range args {
		params = append(params, NewParam(i, NewType(arg.Type), arg.Variadic))
	}

	return params
}

func BuildResults(args []parsing.Argument) []Result {
	var results []Result
	for _, arg := range args {
		results = append(results, NewResult(NewType(arg.Type)))
	}

	return results
}

func BuildBody(receiver Receiver, signature parsing.Signature) []Statement {
	statements := []Statement{
		BuildIncrementStatement(receiver, signature.Name),
	}

	for i, _ := range signature.Params {
		statements = append(statements, BuildAssignStatement(receiver, signature.Name, i, signature.Params))
	}

	if len(signature.Results) > 0 {
		statements = append(statements, BuildReturnStatement(receiver, signature))
	}

	return statements
}

func BuildIncrementStatement(receiver Receiver, name string) IncrementStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", name))
	countField := callField.Type.(Struct).FieldWithName("CallCount")
	selector := NewSelector(receiver, callField, countField)

	return NewIncrementStatement(selector)
}

func BuildAssignStatement(receiver Receiver, name string, index int, args []parsing.Argument) AssignStatement {
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
	param := NewParam(index, NewType(arg.Type), arg.Variadic)

	return NewAssignStatement(selector, param)
}

func BuildReturnStatement(receiver Receiver, signature parsing.Signature) ReturnStatement {
	receiverStruct := receiver.Type.(Pointer).Elem.(NamedType).Type.(Struct)
	callField := receiverStruct.FieldWithName(fmt.Sprintf("%sCall", signature.Name))
	returnsField := callField.Type.(Struct).FieldWithName("Returns")

	var results []Type
	for i, arg := range signature.Results {
		argName := arg.Name
		if argName == "" {
			argName = FieldTypeName(signature.Results, i)
		}

		resultField := returnsField.Type.(Struct).FieldWithName(argName)
		selector := NewSelector(receiver, callField, returnsField, resultField)
		results = append(results, selector)
	}

	return NewReturnStatement(results)
}
