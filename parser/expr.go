package parser

import (
	"fmt"
	"strconv"

	"dreamlang/ast"
	"dreamlang/lexer"
)

// parse_expr 解析表达式。
// 它使用前缀解析函数（NUD）和中缀解析函数（LED）来解析表达式树。
// 参数：
// - p: 指向解析器的指针。
// - bp: 当前绑定的优先级。
// 返回值：
// - 返回解析后的表达式树（ast.Expr）。
// 解析过程：
//  1. 获取当前的 token 类型，并查找对应的 NUD 处理函数。
//  2. 如果找不到对应的 NUD 处理函数，则抛出异常。
//  3. 使用 NUD 处理函数解析当前 token，得到左操作数。
//  4. 当下一个 token 的绑定优先级大于当前优先级时，继续解析：
//     a. 获取当前的 token 类型，并查找对应的 LED 处理函数。
//     b. 如果找不到对应的 LED 处理函数，则抛出异常。
//     c. 使用 LED 处理函数解析当前 token，更新左操作数。
//  5. 返回最终解析得到的表达式树。
func parse_expr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
}

func parse_prefix_expr(p *parser) ast.Expr {
	operatorToken := p.advance()
	expr := parse_expr(p, unary)

	return ast.PrefixExpr{
		Operator: operatorToken,
		Right:    expr,
	}
}

func parse_assignment_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	p.advance()
	rhs := parse_expr(p, bp)

	return ast.AssignmentExpr{
		Assigne:       left,
		AssignedValue: rhs,
	}
}

func parse_range_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	p.advance()

	return ast.RangeExpr{
		Lower: left,
		Upper: parse_expr(p, bp),
	}
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	right := parse_expr(p, defalt_bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Cannot create primary_expr from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parse_member_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	isComputed := p.advance().Kind == lexer.OPEN_BRACKET

	if isComputed {
		rhs := parse_expr(p, bp)
		p.expect(lexer.CLOSE_BRACKET)
		return ast.ComputedExpr{
			Member:   left,
			Property: rhs,
		}
	}

	return ast.MemberExpr{
		Member:   left,
		Property: p.expect(lexer.IDENTIFIER).Value,
	}
}

func parse_array_literal_expr(p *parser) ast.Expr {
	p.expect(lexer.OPEN_BRACKET)
	arrayContents := make([]ast.Expr, 0)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_BRACKET {
		arrayContents = append(arrayContents, parse_expr(p, logical))

		if !p.currentToken().IsOneOfMany(lexer.EOF, lexer.CLOSE_BRACKET) {
			p.expect(lexer.COMMA)
		}
	}

	p.expect(lexer.CLOSE_BRACKET)

	return ast.ArrayLiteral{
		Contents: arrayContents,
	}
}

func parse_grouping_expr(p *parser) ast.Expr {
	p.expect(lexer.OPEN_PAREN)
	expr := parse_expr(p, defalt_bp)
	p.expect(lexer.OPEN_PAREN)
	return expr
}

func parse_call_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	p.advance()
	arguments := make([]ast.Expr, 0)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_PAREN {
		arguments = append(arguments, parse_expr(p, assignment))

		if !p.currentToken().IsOneOfMany(lexer.EOF, lexer.CLOSE_PAREN) {
			p.expect(lexer.COMMA)
		}
	}

	p.expect(lexer.CLOSE_PAREN)
	return ast.CallExpr{
		Method:    left,
		Arguments: arguments,
	}
}

func parse_fn_expr(p *parser) ast.Expr {
	p.expect(lexer.FN)
	functionParams, returnType, functionBody := parse_fn_params_and_body(p)

	return ast.FunctionExpr{
		Parameters: functionParams,
		ReturnType: returnType,
		Body:       functionBody,
	}
}
