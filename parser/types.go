package parser

import (
	"fmt"

	"dreamlang/ast"
	"dreamlang/lexer"
)

type type_nud_handler func(p *parser) ast.Type
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = type_bp_lookup{}
var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}

func type_led(kind lexer.TokenKind, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func type_nud(kind lexer.TokenKind, bp binding_power, nud_fn type_nud_handler) {
	type_bp_lu[kind] = primary
	type_nud_lu[kind] = nud_fn
}

func createTypeTokenLookups() {

	type_nud(lexer.IDENTIFIER, primary, func(p *parser) ast.Type {
		return ast.SymbolType{
			Value: p.advance().Value,
		}
	})

	// []number
	type_nud(lexer.OPEN_BRACKET, member, func(p *parser) ast.Type {
		p.advance()
		p.expect(lexer.CLOSE_BRACKET)
		insideType := parse_type(p, defalt_bp)

		return ast.ListType{
			Underlying: insideType,
		}
	})
}

// parse_type 解析类型表达式并返回解析后的抽象语法树（AST）类型节点。
//
// 参数：
// - p: 指向解析器实例的指针，用于访问当前的令牌和解析状态。
// - bp: 当前绑定的优先级，用于控制解析的优先级和结合性。
//
// 返回值：
// - 返回解析后的抽象语法树（AST）类型节点。
//
// 该函数首先根据当前令牌类型查找相应的NUD（null denotation）处理函数，
// 如果找不到对应的处理函数则抛出异常。然后调用NUD处理函数解析当前令牌。
// 接着在一个循环中，根据当前令牌类型和绑定优先级查找相应的LED（left denotation）处理函数，
// 如果找不到对应的处理函数则抛出异常。调用LED处理函数解析后续的令牌，
// 直到当前令牌的优先级不再大于传入的绑定优先级为止。
// 最终返回解析后的AST类型节点。
func parse_type(p *parser, bp binding_power) ast.Type {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := type_nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("type: NUD Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for type_bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := type_led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("type: LED Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
}
