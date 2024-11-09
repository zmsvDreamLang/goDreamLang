package parser

import (
	"fmt"

	"dreamlang/lexer"
)

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) nextToken() lexer.Token {
	return p.tokens[p.pos+1]
}

func (p *parser) previousToken() lexer.Token {
	return p.tokens[p.pos-1]
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.tokens[p.pos].Kind
}

// expectError 检查当前 token 是否为预期的类型，如果不是则触发错误。
//
// 参数:
//   - expectedKind: 预期的 token 类型。
//   - err: 如果 token 类型不匹配时触发的错误信息。如果为 nil，则会生成一个默认的错误信息。
//
// 返回值:
//   - lexer.Token: 如果当前 token 类型匹配预期类型，则返回当前 token 并将解析器推进到下一个 token。
//
// 错误:
//   - 如果当前 token 类型不匹配预期类型，则触发 panic 并输出错误信息。
func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected %s but recieved %s instead\n", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}

	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
