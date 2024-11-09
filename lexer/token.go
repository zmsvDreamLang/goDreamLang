package lexer

import "fmt"

type TokenKind int

const (
	TokenTypeEOF TokenKind = iota
	TokenTypeValNull
	TokenTypeValTrue
	TokenTypeValFalse
	TokenTypeValNumber
	TokenTypeValString
	TokenTypeValIdentifier

	// Grouping & Braces
	TokenTypeSymbolLBracket
	TokenTypeSymbolRBracket
	TokenTypeSymbolLBrance
	TokenTypeSymbolRBrance
	TokenTypeSymbolLParen
	TokenTypeSymbolRParen

	// Equivilance
	TokenTypeSymbolAssignment
	TokenTypeSymbolEqual
	TokenTypeSymbolNotEqual

	// Conditional
	TokenTypeSymbolLT
	TokenTypeSymbolLTEQ
	TokenTypeSymbolGT
	TokenTypeSymbolGTEQ

	// Logical
	TokenTypeSymbolOr
	TokenTypeSymbolAnd
	TokenTypeSymbolNot
	TokenTypeSymbolXor
	TokenTypeSymbolXorNot
	TokenTypeSymbolBitOr
	TokenTypeSymbolBitAnd
	TokenTypeSymbolBitNot
	TokenTypeSymbolBitXor
	TokenTypeSymbolBitXorNot

	// Symbols
	TokenTypeSymbolDot
	TokenTypeSymbolConcat
	TokenTypeSymbolVarargs
	TokenTypeSymbolSemiColon
	TokenTypeSymbolColon
	TokenTypeSymbolQuestion
	TokenTypeSymbolComma

	// Shorthand
	TokenTypeSymbolPlusPlus
	TokenTypeSymbolMinusMinus
	TokenTypeSymbolPlusEqual
	TokenTypeSymbolDashEqual
	TokenTypeSymbolStarEqual
	TokenTypeSymbolSlashEqual
	TokenTypeSymbolPercentEqual

	//Maths
	TokenTypeSymbolPlus
	TokenTypeSymbolDash
	TokenTypeSymbolStar
	TokenTypeSymbolSlash
	TokenTypeSymbolPercent

	// Reserved Keywords
	TokenTypeKeywordVar
	TokenTypeKeywordLet
	TokenTypeKeywordVal
	TokenTypeKeywordClass
	TokenTypeKeywordNew
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN

	// Misc
	NUM_TOKENS
)

var reserved_lu map[string]TokenKind = map[string]TokenKind{
	"true":    TRUE,
	"false":   FALSE,
	"null":    NULL,
	"let":     LET,
	"const":   CONST,
	"class":   CLASS,
	"new":     NEW,
	"import":  IMPORT,
	"from":    FROM,
	"fn":      FN,
	"if":      IF,
	"else":    ELSE,
	"foreach": FOREACH,
	"while":   WHILE,
	"for":     FOR,
	"export":  EXPORT,
	"typeof":  TYPEOF,
	"in":      IN,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (tk Token) IsOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == tk.Kind {
			return true
		}
	}

	return false
}

func (token Token) Debug() {
	if token.Kind == IDENTIFIER || token.Kind == NUMBER || token.Kind == STRING {
		fmt.Printf("%s(%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s()\n", TokenKindString(token.Kind))
	}
}

// TokenKindString 根据给定的 TokenKind 返回对应的字符串表示。
//
// 参数:
//   - kind: TokenKind 类型的值，表示要转换的标记类型。
//
// 返回值:
//   - string: 对应于给定 TokenKind 的字符串表示。如果 TokenKind 未知，则返回 "unknown" 加上其数值表示。
//
// 支持的 TokenKind 类型及其对应的字符串表示包括但不限于:
//   - EOF: "eof"
//   - NULL: "null"
//   - NUMBER: "number"
//   - STRING: "string"
//   - TRUE: "true"
//   - FALSE: "false"
//   - IDENTIFIER: "identifier"
//   - OPEN_BRACKET: "open_bracket"
//   - CLOSE_BRACKET: "close_bracket"
//   - OPEN_CURLY: "open_curly"
//   - CLOSE_CURLY: "close_curly"
//   - OPEN_PAREN: "open_paren"
//   - CLOSE_PAREN: "close_paren"
//   - ASSIGNMENT: "assignment"
//   - EQUALS: "equals"
//   - NOT_EQUALS: "not_equals"
//   - NOT: "not"
//   - LESS: "less"
//   - LESS_EQUALS: "less_equals"
//   - GREATER: "greater"
//   - GREATER_EQUALS: "greater_equals"
//   - OR: "or"
//   - AND: "and"
//   - DOT: "dot"
//   - DOT_DOT: "dot_dot"
//   - SEMI_COLON: "semi_colon"
//   - COLON: "colon"
//   - QUESTION: "question"
//   - COMMA: "comma"
//   - PLUS_PLUS: "plus_plus"
//   - MINUS_MINUS: "minus_minus"
//   - PLUS_EQUALS: "plus_equals"
//   - MINUS_EQUALS: "minus_equals"
//   - NULLISH_ASSIGNMENT: "nullish_assignment"
//   - PLUS: "plus"
//   - DASH: "dash"
//   - SLASH: "slash"
//   - STAR: "star"
//   - PERCENT: "percent"
//   - LET: "let"
//   - CONST: "const"
//   - CLASS: "class"
//   - NEW: "new"
//   - IMPORT: "import"
//   - FROM: "from"
//   - FN: "fn"
//   - IF: "if"
//   - ELSE: "else"
//   - FOREACH: "foreach"
//   - FOR: "for"
//   - WHILE: "while"
//   - EXPORT: "export"
//   - IN: "in"
func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NULL:
		return "null"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSIGNMENT:
		return "assignment"
	case EQUALS:
		return "equals"
	case NOT_EQUALS:
		return "not_equals"
	case NOT:
		return "not"
	case LESS:
		return "less"
	case LESS_EQUALS:
		return "less_equals"
	case GREATER:
		return "greater"
	case GREATER_EQUALS:
		return "greater_equals"
	case OR:
		return "or"
	case AND:
		return "and"
	case DOT:
		return "dot"
	case DOT_DOT:
		return "dot_dot"
	case SEMI_COLON:
		return "semi_colon"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case NULLISH_ASSIGNMENT:
		return "nullish_assignment"
	case PLUS:
		return "plus"
	case DASH:
		return "dash"
	case SLASH:
		return "slash"
	case STAR:
		return "star"
	case PERCENT:
		return "percent"
	case LET:
		return "let"
	case CONST:
		return "const"
	case CLASS:
		return "class"
	case NEW:
		return "new"
	case IMPORT:
		return "import"
	case FROM:
		return "from"
	case FN:
		return "fn"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOREACH:
		return "foreach"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case EXPORT:
		return "export"
	case IN:
		return "in"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}

func newUniqueToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
}
