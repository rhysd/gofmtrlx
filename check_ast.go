package main

import (
	"fmt"
	"go/ast"
)

type badNodeDetector struct {
	found ast.Node
}

func (v *badNodeDetector) Visit(node ast.Node) ast.Visitor {
	_, ok := node.(*ast.BadExpr)
	if ok {
		v.found = node
		return nil
	}
	_, ok = node.(*ast.BadStmt)
	if ok {
		v.found = node
		return nil
	}
	_, ok = node.(*ast.BadDecl)
	if ok {
		v.found = node
		return nil
	}

	return v
}

func checkBadAST(f *ast.File, originalError error) error {
	v := &badNodeDetector{nil}
	ast.Walk(v, f)
	if v.found != nil {
		fmt.Printf("%v", v.found)
		return originalError
	}
	return nil
}
