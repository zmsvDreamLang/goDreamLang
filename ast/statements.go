package ast

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

// VarDeclarationStmt 表示一个变量声明语句。
//
// 字段:
// - Identifier: 变量的标识符名称。
// - Constant: 一个布尔值，指示变量是否为常量。如果为 true，则表示该变量为常量。
// - AssignedValue: 变量的初始赋值表达式。
// - ExplicitType: 变量的显式类型。如果未指定类型，则可能为 nil。
type VarDeclarationStmt struct {
	Identifier    string
	Constant      bool
	AssignedValue Expr
	ExplicitType  Type
}

func (n VarDeclarationStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt() {}

type Parameter struct {
	Name string
	Type Type
}

type FunctionDeclarationStmt struct {
	Parameters []Parameter
	Name       string
	Body       []Stmt
	ReturnType Type
}

func (n FunctionDeclarationStmt) stmt() {}

type IfStmt struct {
	Condition  Expr
	Consequent Stmt
	Alternate  Stmt
}

func (n IfStmt) stmt() {}

type ImportStmt struct {
	Name string
	From string
}

func (n ImportStmt) stmt() {}

type ForeachStmt struct {
	Value    string
	Index    bool
	Iterable Expr
	Body     []Stmt
}

func (n ForeachStmt) stmt() {}

type ClassDeclarationStmt struct {
	Name string
	Body []Stmt
}

func (n ClassDeclarationStmt) stmt() {}
