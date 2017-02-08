// Copyright 2017 rhysd <lin90162@yahoo.co.jp>
//
// This file is not a part of gofmt.
// https://github.com/rhysd/gofmtrlx

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go/ast"
	"go/scanner"
	"os"
)

type badNodeDetector struct {
	reason string
}

func (v *badNodeDetector) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.BadExpr:
		v.reason = fmt.Sprintf("Bad expression around offset %d", node.Pos())
		return nil
	case *ast.BadStmt:
		v.reason = fmt.Sprintf("Bad statement around offset %d", node.Pos())
		return nil
	case *ast.BadDecl:
		v.reason = fmt.Sprintf("Bad declaration around offset %d", node.Pos())
		return nil
	}
	return v
}

func checkBadAST(f *ast.File, originalError error) error {
	v := &badNodeDetector{""}
	ast.Walk(v, f)
	if v.reason != "" {
		return errors.Wrapf(originalError, "Cannot format code because of bad node: %s", v.reason)
	}
	scanner.PrintError(os.Stderr, originalError)
	return nil
}
