package parser

import (
	"fmt"

	"dreamlang/ast"
	"dreamlang/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	return parse_expression_stmt(p)
}

func parse_expression_stmt(p *parser) ast.ExpressionStmt {
	expression := parse_expr(p, defalt_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_block_stmt(p *parser) ast.Stmt {
	p.expect(lexer.OPEN_CURLY)
	body := []ast.Stmt{}

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		body = append(body, parse_stmt(p))
	}

	p.expect(lexer.CLOSE_CURLY)
	return ast.BlockStmt{
		Body: body,
	}
}

// parse_var_decl_stmt 解析变量声明语句，并返回一个 ast.Stmt 类型的节点。
//
// 该函数处理以下几种情况：
// 1. 常量声明：以 `const` 关键字开头。
// 2. 变量声明：以 `var` 关键字开头。
//
// 参数：
// - p: *parser 类型的指针，用于解析输入的 token。
//
// 返回值：
// - ast.Stmt: 表示变量声明的语法树节点。
//
// 解析过程：
// 1. 获取当前 token 的类型，判断是否为常量声明。
// 2. 期望下一个 token 为标识符（变量名），否则抛出错误。
// 3. 如果下一个 token 为冒号（`:`），则解析变量的显式类型。
// 4. 如果下一个 token 不是分号（`;`），则期望为赋值操作符（`=`），并解析赋值表达式。
// 5. 如果没有显式类型且没有赋值表达式，则抛出错误。
// 6. 期望下一个 token 为分号（`;`），表示声明语句结束。
// 7. 如果是常量声明但没有赋值表达式，则抛出错误。
//
// 返回的 ast.VarDeclarationStmt 包含以下字段：
// - Constant: 是否为常量声明。
// - Identifier: 变量名。
// - AssignedValue: 赋值表达式（如果有的话）。
// - ExplicitType: 显式类型（如果有的话）。
func parse_var_decl_stmt(p *parser) ast.Stmt {
	var explicitType ast.Type
	startToken := p.advance().Kind
	isConstant := startToken == lexer.CONST
	symbolName := p.expectError(lexer.IDENTIFIER,
		fmt.Sprintf("Following %s expected variable name however instead recieved %s instead\n",
			lexer.TokenKindString(startToken), lexer.TokenKindString(p.currentTokenKind())))

	if p.currentTokenKind() == lexer.COLON {
		p.expect(lexer.COLON)
		explicitType = parse_type(p, defalt_bp)
	}

	var assignmentValue ast.Expr
	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		assignmentValue = parse_expr(p, assignment)
	} else if explicitType == nil {
		panic("Missing explicit type for variable declaration.")
	}

	p.expect(lexer.SEMI_COLON)

	if isConstant && assignmentValue == nil {
		panic("Cannot define constant variable without providing default value.")
	}

	return ast.VarDeclarationStmt{
		Constant:      isConstant,
		Identifier:    symbolName.Value,
		AssignedValue: assignmentValue,
		ExplicitType:  explicitType,
	}
}

func parse_fn_params_and_body(p *parser) ([]ast.Parameter, ast.Type, []ast.Stmt) {
	functionParams := make([]ast.Parameter, 0)

	p.expect(lexer.OPEN_PAREN)
	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_PAREN {
		paramName := p.expect(lexer.IDENTIFIER).Value
		p.expect(lexer.COLON)
		paramType := parse_type(p, defalt_bp)

		functionParams = append(functionParams, ast.Parameter{
			Name: paramName,
			Type: paramType,
		})

		if !p.currentToken().IsOneOfMany(lexer.CLOSE_PAREN, lexer.EOF) {
			p.expect(lexer.COMMA)
		}
	}

	p.expect(lexer.CLOSE_PAREN)
	var returnType ast.Type

	if p.currentTokenKind() == lexer.COLON {
		p.advance()
		returnType = parse_type(p, defalt_bp)
	}

	functionBody := ast.ExpectStmt[ast.BlockStmt](parse_block_stmt(p)).Body

	return functionParams, returnType, functionBody
}

func parse_fn_declaration(p *parser) ast.Stmt {
	p.advance()
	functionName := p.expect(lexer.IDENTIFIER).Value
	functionParams, returnType, functionBody := parse_fn_params_and_body(p)

	return ast.FunctionDeclarationStmt{
		Parameters: functionParams,
		ReturnType: returnType,
		Body:       functionBody,
		Name:       functionName,
	}
}

func parse_if_stmt(p *parser) ast.Stmt {
	p.advance()
	condition := parse_expr(p, assignment)
	consequent := parse_block_stmt(p)

	var alternate ast.Stmt
	if p.currentTokenKind() == lexer.ELSE {
		p.advance()

		if p.currentTokenKind() == lexer.IF {
			alternate = parse_if_stmt(p)
		} else {
			alternate = parse_block_stmt(p)
		}
	}

	return ast.IfStmt{
		Condition:  condition,
		Consequent: consequent,
		Alternate:  alternate,
	}
}

func parse_import_stmt(p *parser) ast.Stmt {
	p.advance()
	var importFrom string
	importName := p.expect(lexer.IDENTIFIER).Value

	if p.currentTokenKind() == lexer.FROM {
		p.advance()
		importFrom = p.expect(lexer.STRING).Value
	} else {
		importFrom = importName
	}

	p.expect(lexer.SEMI_COLON)
	return ast.ImportStmt{
		Name: importName,
		From: importFrom,
	}
}

func parse_foreach_stmt(p *parser) ast.Stmt {
	p.advance()
	valueName := p.expect(lexer.IDENTIFIER).Value

	var index bool
	if p.currentTokenKind() == lexer.COMMA {
		p.expect(lexer.COMMA)
		p.expect(lexer.IDENTIFIER)
		index = true
	}

	p.expect(lexer.IN)
	iterable := parse_expr(p, defalt_bp)
	body := ast.ExpectStmt[ast.BlockStmt](parse_block_stmt(p)).Body

	return ast.ForeachStmt{
		Value:    valueName,
		Index:    index,
		Iterable: iterable,
		Body:     body,
	}
}

func parse_class_declaration_stmt(p *parser) ast.Stmt {
	p.advance()
	className := p.expect(lexer.IDENTIFIER).Value
	classBody := parse_block_stmt(p)

	return ast.ClassDeclarationStmt{
		Name: className,
		Body: ast.ExpectStmt[ast.BlockStmt](classBody).Body,
	}
}
