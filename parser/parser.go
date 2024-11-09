package parser

import (
	"dreamlang/ast"
	"dreamlang/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	createTypeTokenLookups()

	p := &parser{
		tokens: tokens,
		pos:    0,
	}

	return p
}

// Parse 函数解析给定的源代码字符串，并返回一个表示代码块的 ast.BlockStmt 结构体。
//
// 参数:
// - source: 一个包含源代码的字符串。
//
// 返回值:
// - ast.BlockStmt: 一个包含解析后的语句列表的代码块结构体。
//
// 该函数首先使用 lexer.Tokenize 函数将源代码字符串转换为标记列表。
// 然后，它创建一个解析器实例，并初始化一个空的语句列表。
// 在解析器还有标记可供解析时，它会不断调用 parse_stmt 函数解析语句，并将解析后的语句添加到语句列表中。
// 最后，函数返回一个包含所有解析语句的 ast.BlockStmt 结构体。
func Parse(source string) ast.BlockStmt {
	tokens := lexer.Tokenize(source)
	p := createParser(tokens)
	body := make([]ast.Stmt, 0)

	for p.hasTokens() {
		body = append(body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}
