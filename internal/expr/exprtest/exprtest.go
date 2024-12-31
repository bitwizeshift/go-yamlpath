/*
Package exprtest provides test-doubles for [expr.Expr] implementations.
*/
package exprtest

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
)

// Func is a function that implements the Expr interface
type Func func(*expr.Context) ([]*yaml.Node, error)

// Eval calls the function
func (f Func) Eval(ctx *expr.Context) ([]*yaml.Node, error) {
	return f(ctx)
}

var _ expr.Expr = (*Func)(nil)

// Empty returns an Expr that always returns an empty result.
func Empty() expr.Expr {
	return Func(func(*expr.Context) ([]*yaml.Node, error) {
		return nil, nil
	})
}

// Error returns an Expr that always returns an error.
func Error(err error) expr.Expr {
	return Func(func(*expr.Context) ([]*yaml.Node, error) {
		return nil, err
	})
}

// Return returns an Expr that returns the given nodes.
func Return(nodes ...*yaml.Node) expr.Expr {
	return Func(func(*expr.Context) ([]*yaml.Node, error) {
		return nodes, nil
	})
}
