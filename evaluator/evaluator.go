package evaluator

import (
	"github.com/FumiKimura/interpreter-with-golang/ast"
	"github.com/FumiKimura/interpreter-with-golang/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}
