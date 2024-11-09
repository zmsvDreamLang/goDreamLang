package ast

import "dreamlang/helpers"

type Stmt interface {
	stmt()
}

type Expr interface {
	expr()
}

type Type interface {
	_type()
}

// ExpectExpr 函数接受一个表达式，并期望该表达式是特定类型 T。
// 该函数使用了泛型，允许在调用时指定具体的表达式类型。
//
// 参数:
// - expr: 一个实现了 Expr 接口的表达式。
//
// 返回值:
// - 返回类型为 T 的表达式。
//
// 示例:
// ```go
// var myExpr Expr = ...
// var result = ExpectExpr[SpecificExprType](myExpr)
// ```
func ExpectExpr[T Expr](expr Expr) T {
	return helpers.ExpectType[T](expr)
}

func ExpectStmt[T Stmt](expr Stmt) T {
	return helpers.ExpectType[T](expr)
}
