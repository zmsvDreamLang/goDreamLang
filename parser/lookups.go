package parser

import (
	"dreamlang/ast"
	"dreamlang/lexer"
)

type binding_power int

const (
	defalt_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu = stmt_lookup{}

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, bp binding_power, nud_fn nud_handler) {
	bp_lu[kind] = primary
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = defalt_bp
	stmt_lu[kind] = stmt_fn
}

// createTokenLookups 函数用于创建和初始化各种令牌的查找表。
// 这些查找表用于解析不同类型的表达式和语句。
// 具体包括以下几类：
//
// 1. 赋值操作符：
//   - lexer.ASSIGNMENT
//   - lexer.PLUS_EQUALS
//   - lexer.MINUS_EQUALS
//
// 2. 逻辑操作符：
//   - lexer.AND
//   - lexer.OR
//   - lexer.DOT_DOT
//
// 3. 关系操作符：
//   - lexer.LESS
//   - lexer.LESS_EQUALS
//   - lexer.GREATER
//   - lexer.GREATER_EQUALS
//   - lexer.EQUALS
//   - lexer.NOT_EQUALS
//
// 4. 加法和乘法操作符：
//   - lexer.PLUS
//   - lexer.DASH
//   - lexer.SLASH
//   - lexer.STAR
//   - lexer.PERCENT
//
// 5. 字面量和符号：
//   - lexer.NUMBER
//   - lexer.STRING
//   - lexer.IDENTIFIER
//
// 6. 一元/前缀操作符：
//   - lexer.TYPEOF
//   - lexer.DASH
//   - lexer.NOT
//   - lexer.OPEN_BRACKET
//
// 7. 成员/计算/调用操作符：
//   - lexer.DOT
//   - lexer.OPEN_BRACKET
//   - lexer.OPEN_PAREN
//
// 8. 分组表达式：
//   - lexer.OPEN_PAREN
//   - lexer.FN
//   - lexer.NEW
//
// 9. 语句：
//   - lexer.OPEN_CURLY
//   - lexer.LET
//   - lexer.CONST
//   - lexer.FN
//   - lexer.IF
//   - lexer.IMPORT
//   - lexer.FOREACH
//   - lexer.CLASS
//
// 该函数通过调用 led、nud 和 stmt 函数来为每种令牌类型注册相应的解析函数。
func createTokenLookups() {
	// Assignment
	led(lexer.ASSIGNMENT, assignment, parse_assignment_expr)
	led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)

	// Logical
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)
	led(lexer.DOT_DOT, logical, parse_range_expr)

	// Relational
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	led(lexer.EQUALS, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)

	// Additive & Multiplicitave
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)

	// Literals & Symbols
	nud(lexer.NUMBER, primary, parse_primary_expr)
	nud(lexer.STRING, primary, parse_primary_expr)
	nud(lexer.IDENTIFIER, primary, parse_primary_expr)

	// Unary/Prefix
	nud(lexer.TYPEOF, unary, parse_prefix_expr)
	nud(lexer.DASH, unary, parse_prefix_expr)
	nud(lexer.NOT, unary, parse_prefix_expr)
	nud(lexer.OPEN_BRACKET, primary, parse_array_literal_expr)

	// Member / Computed // Call
	led(lexer.DOT, member, parse_member_expr)
	led(lexer.OPEN_BRACKET, member, parse_member_expr)
	led(lexer.OPEN_PAREN, call, parse_call_expr)

	// Grouping Expr
	nud(lexer.OPEN_PAREN, defalt_bp, parse_grouping_expr)
	nud(lexer.FN, defalt_bp, parse_fn_expr)
	nud(lexer.NEW, defalt_bp, func(p *parser) ast.Expr {
		p.advance()
		classInstantiation := parse_expr(p, defalt_bp)

		return ast.NewExpr{
			Instantiation: ast.ExpectExpr[ast.CallExpr](classInstantiation),
		}
	})

	stmt(lexer.OPEN_CURLY, parse_block_stmt)
	stmt(lexer.LET, parse_var_decl_stmt)
	stmt(lexer.CONST, parse_var_decl_stmt)
	stmt(lexer.FN, parse_fn_declaration)
	stmt(lexer.IF, parse_if_stmt)
	stmt(lexer.IMPORT, parse_import_stmt)
	stmt(lexer.FOREACH, parse_foreach_stmt)
	stmt(lexer.CLASS, parse_class_declaration_stmt)
}
