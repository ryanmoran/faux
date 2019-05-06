package rendering

import "go/ast"

type Func struct {
	Name     string
	Receiver Receiver
	Params   []Param
	Results  []Result
	Body     []Statement
}

func NewFunc(name string, receiver Receiver, params []Param, results []Result, body []Statement) Func {
	return Func{
		Name:     name,
		Receiver: receiver,
		Params:   params,
		Results:  results,
		Body:     body,
	}
}

func (f Func) Expr() ast.Expr {
	var params []*ast.Field
	for _, param := range f.Params {
		params = append(params, param.Field())
	}

	var results []*ast.Field
	for _, result := range f.Results {
		results = append(results, result.Field())
	}

	return &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: params,
			},
			Results: &ast.FieldList{
				List: results,
			},
		},
	}
}

func (f Func) isType() {}

func (f Func) Decl() ast.Decl {
	var params []*ast.Field
	for _, param := range f.Params {
		params = append(params, param.Field())
	}

	var results []*ast.Field
	for _, result := range f.Results {
		results = append(results, result.Field())
	}

	var statements []ast.Stmt
	for _, statement := range f.Body {
		statements = append(statements, statement.Stmt())
	}

	return &ast.FuncDecl{
		Name: f.Ident(),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				f.Receiver.Field(),
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: params,
			},
			Results: &ast.FieldList{
				List: results,
			},
		},
		Body: &ast.BlockStmt{
			List: statements,
		},
	}
}

func (f Func) Ident() *ast.Ident {
	return ast.NewIdent(f.Name)
}
