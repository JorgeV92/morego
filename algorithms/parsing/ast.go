package parsing

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
)

type FunctionInfo struct {
	Name           string
	Receiver       string
	ParameterCount int
	ResultCount    int
	Calls          []string
}

func ParseFunctionInfo(src string) ([]FunctionInfo, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "sample.go", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	functions := make([]FunctionInfo, 0)
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		info := FunctionInfo{
			Name:           fn.Name.Name,
			Receiver:       receiverString(fn.Recv),
			ParameterCount: fieldCount(fn.Type.Params),
			ResultCount:    fieldCount(fn.Type.Results),
		}

		if fn.Body != nil {
			ast.Inspect(fn.Body, func(node ast.Node) bool {
				call, ok := node.(*ast.CallExpr)
				if !ok {
					return true
				}

				info.Calls = append(info.Calls, nodeString(call.Fun))
				return true
			})
		}

		functions = append(functions, info)
	}

	return functions, nil
}

func fieldCount(fields *ast.FieldList) int {
	if fields == nil {
		return 0
	}

	count := 0
	for _, field := range fields.List {
		if len(field.Names) == 0 {
			count++
			continue
		}
		count += len(field.Names)
	}

	return count
}

func receiverString(recv *ast.FieldList) string {
	if recv == nil || len(recv.List) == 0 {
		return ""
	}

	return nodeString(recv.List[0].Type)
}

func nodeString(node ast.Node) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, token.NewFileSet(), node); err != nil {
		return ""
	}
	return buf.String()
}
