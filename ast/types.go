package ast

type SymbolType struct {
	Value string
}

func (t SymbolType) _type() {}

// ListType 表示一个列表类型。
// 它包含一个基础类型（Underlying），该基础类型定义了列表中元素的类型。
type ListType struct {
	Underlying Type
}

func (t ListType) _type() {}
