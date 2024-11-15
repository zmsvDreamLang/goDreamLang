package ast

import (
	"dreamlang/lexer"
)

// --------------------
// Literal Expressions
// --------------------

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// --------------------
// Complex Expressions
// --------------------

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {}

type AssignmentExpr struct {
	Assigne       Expr
	AssignedValue Expr
}

func (n AssignmentExpr) expr() {}

type PrefixExpr struct {
	Operator lexer.Token
	Right    Expr
}

func (n PrefixExpr) expr() {}

type MemberExpr struct {
	Member   Expr
	Property string
}

func (n MemberExpr) expr() {}

type CallExpr struct {
	Method    Expr
	Arguments []Expr
}

func (n CallExpr) expr() {}

type ComputedExpr struct {
	Member   Expr
	Property Expr
}

func (n ComputedExpr) expr() {}

type RangeExpr struct {
	Lower Expr
	Upper Expr
}

func (n RangeExpr) expr() {}

// FunctionExpr 表示一个函数表达式。
//
// 字段:
// - Parameters: 函数的参数列表，包含多个 Parameter 类型的元素。
// - Body: 函数的主体，由多个 Stmt 类型的元素组成，表示函数的执行语句。
// - ReturnType: 函数的返回类型，表示函数执行后返回的值的类型。
type FunctionExpr struct {
	Parameters []Parameter
	Body       []Stmt
	ReturnType Type
}

func (n FunctionExpr) expr() {}

type ArrayLiteral struct {
	Contents []Expr
}

func (n ArrayLiteral) expr() {}

type NewExpr struct {
	Instantiation CallExpr
}

func (n NewExpr) expr() {}
