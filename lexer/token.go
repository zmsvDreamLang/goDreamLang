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
	TokenTypeSymbolLShift
	TokenTypeSymbolRShift
	TokenTypeSymbolLArrow
	TokenTypeSymbolRArrow

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
	TokenTypeKeywordPublic
	TokenTypeKeywordPrivate
	TokenTypeKeywordStatic
	TokenTypeKeywordFinal
	TokenTypeKeywordAbstract
	TokenTypeKeywordOverride
	TokenTypeKeywordImport
	TokenTypeKeywordFrom
	TokenTypeKeywordFunc
	TokenTypeKeywordIf
	TokenTypeKeywordElse
	TokenTypeKeywordElseIf
	TokenTypeKeywordSwitch
	TokenTypeKeywordCase
	TokenTypeKeywordDefault
	TokenTypeKeywordForeach
	TokenTypeKeywordIn
	TokenTypeKeywordFor
	TokenTypeKeywordWhile
	TokenTypeKeywordExport

	// Misc
	NUM_TOKENS
)

var reserved_lu map[string]TokenKind = map[string]TokenKind{
	"var":      TokenTypeKeywordVar,
	"let":      TokenTypeKeywordLet,
	"val":      TokenTypeKeywordVal,
	"class":    TokenTypeKeywordClass,
	"new":      TokenTypeKeywordNew,
	"public":   TokenTypeKeywordPublic,
	"private":  TokenTypeKeywordPrivate,
	"static":   TokenTypeKeywordStatic,
	"final":    TokenTypeKeywordFinal,
	"abstract": TokenTypeKeywordAbstract,
	"override": TokenTypeKeywordOverride,
	"import":   TokenTypeKeywordImport,
	"from":     TokenTypeKeywordFrom,
	"func":     TokenTypeKeywordFunc,
	"if":       TokenTypeKeywordIf,
	"else":     TokenTypeKeywordElse,
	"elseif":   TokenTypeKeywordElseIf,
	"switch":   TokenTypeKeywordSwitch,
	"case":     TokenTypeKeywordCase,
	"default":  TokenTypeKeywordDefault,
	"foreach":  TokenTypeKeywordForeach,
	"in":       TokenTypeKeywordIn,
	"for":      TokenTypeKeywordFor,
	"while":    TokenTypeKeywordWhile,
	"export":   TokenTypeKeywordExport,
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
	if token.Kind == TokenTypeValIdentifier || token.Kind == TokenTypeValNumber || token.Kind == TokenTypeValString {
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
//  - TokenTypeEOF: "eof"
//  - TokenTypeValNull: "null"
//  - TokenTypeValTrue: "true"
//  - TokenTypeValFalse: "false"
//  - TokenTypeValNumber: "number"
//  - TokenTypeValString: "string"
//  - TokenTypeValIdentifier: "identifier"
//  - TokenTypeSymbolLBracket: "["
//  - TokenTypeSymbolRBracket: "]"
//  - TokenTypeSymbolLBrance: "{"
//  - TokenTypeSymbolRBrance: "}"
//  - TokenTypeSymbolLParen: "("
//  - TokenTypeSymbolRParen: ")"
//  - TokenTypeSymbolAssignment: "="
//  - TokenTypeSymbolEqual: "=="
//  - TokenTypeSymbolNotEqual: "!="
//  - TokenTypeSymbolLT: "<"
//  - TokenTypeSymbolLTEQ: "<="
//  - TokenTypeSymbolGT: ">"
//  - TokenTypeSymbolGTEQ: ">="
//  - TokenTypeSymbolOr: "||"
//  - TokenTypeSymbolAnd: "&&"
//  - TokenTypeSymbolNot: "!"
//  - TokenTypeSymbolXor: "^^"
//  - TokenTypeSymbolXorNot: "^!"
//  - TokenTypeSymbolBitOr: "|"
//  - TokenTypeSymbolBitAnd: "&"
//  - TokenTypeSymbolBitNot: "~"
//  - TokenTypeSymbolBitXor: "^"
//  - TokenTypeSymbolBitXorNot: "^~"
//  - TokenTypeSymbolLShift: "<<"
//  - TokenTypeSymbolRShift: ">>"
//  - TokenTypeSymbolLArrow: "<-"
//  - TokenTypeSymbolRArrow: "->"
//  - TokenTypeSymbolDot: "."
//  - TokenTypeSymbolConcat: ".."
//  - TokenTypeSymbolVarargs: "..."
//  - TokenTypeSymbolSemiColon: ";"
//  - TokenTypeSymbolColon: ":"
//  - TokenTypeSymbolQuestion: "?"
//  - TokenTypeSymbolComma: ","
//  - TokenTypeSymbolPlusPlus: "++"
//  - TokenTypeSymbolMinusMinus: "--"
//  - TokenTypeSymbolPlusEqual: "+="
//  - TokenTypeSymbolDashEqual: "-="
//  - TokenTypeSymbolStarEqual: "*="
//  - TokenTypeSymbolSlashEqual: "/="
//  - TokenTypeSymbolPercentEqual: "%="
//  - TokenTypeSymbolPlus: "+"
//  - TokenTypeSymbolDash: "-"
//  - TokenTypeSymbolStar: "*"
//  - TokenTypeSymbolSlash: "/"
//  - TokenTypeSymbolPercent: "%"
//  - TokenTypeKeywordVar: "var"
//  - TokenTypeKeywordLet: "let"
//  - TokenTypeKeywordVal: "val"
//  - TokenTypeKeywordClass: "class"
//  - TokenTypeKeywordNew: "new"
//  - TokenTypeKeywordImport: "import"
//  - TokenTypeKeywordFrom: "from"
//  - TokenTypeKeywordFunc: "func"
//  - TokenTypeKeywordIf: "if"
//  - TokenTypeKeywordElse: "else"
//  - TokenTypeKeywordElseIf: "elseif"
//  - TokenTypeKeywordSwitch: "switch"
//  - TokenTypeKeywordCase: "case"
//  - TokenTypeKeywordDefault: "default"
//  - TokenTypeKeywordForeach: "foreach"
//  - TokenTypeKeywordFor: "for"
//  - TokenTypeKeywordWhile: "while"
//  - TokenTypeKeywordExport: "export"
//  - TokenTypeKeywordIn: "in"

func TokenKindString(kind TokenKind) string {
	switch kind {
	case TokenTypeEOF:
		return "eof"
	case TokenTypeValNull:
		return "null"
	case TokenTypeValTrue:
		return "true"
	case TokenTypeValFalse:
		return "false"
	case TokenTypeValNumber:
		return "number"
	case TokenTypeValString:
		return "string"
	case TokenTypeValIdentifier:
		return "identifier"
	case TokenTypeSymbolLBracket:
		return "["
	case TokenTypeSymbolRBracket:
		return "]"
	case TokenTypeSymbolLBrance:
		return "{"
	case TokenTypeSymbolRBrance:
		return "}"
	case TokenTypeSymbolLParen:
		return "("
	case TokenTypeSymbolRParen:
		return ")"
	case TokenTypeSymbolAssignment:
		return "="
	case TokenTypeSymbolEqual:
		return "=="
	case TokenTypeSymbolNotEqual:
		return "!="
	case TokenTypeSymbolLT:
		return "<"
	case TokenTypeSymbolLTEQ:
		return "<="
	case TokenTypeSymbolGT:
		return ">"
	case TokenTypeSymbolGTEQ:
		return ">="
	case TokenTypeSymbolOr:
		return "||"
	case TokenTypeSymbolAnd:
		return "&&"
	case TokenTypeSymbolNot:
		return "!"
	case TokenTypeSymbolXor:
		return "^^"
	case TokenTypeSymbolXorNot:
		return "^!"
	case TokenTypeSymbolBitOr:
		return "|"
	case TokenTypeSymbolBitAnd:
		return "&"
	case TokenTypeSymbolBitNot:
		return "~"
	case TokenTypeSymbolBitXor:
		return "^"
	case TokenTypeSymbolBitXorNot:
		return "^~"
	case TokenTypeSymbolLShift:
		return "<<"
	case TokenTypeSymbolRShift:
		return ">>"
	case TokenTypeSymbolLArrow:
		return "<-"
	case TokenTypeSymbolRArrow:
		return "->"
	case TokenTypeSymbolDot:
		return "."
	case TokenTypeSymbolConcat:
		return ".."
	case TokenTypeSymbolVarargs:
		return "..."
	case TokenTypeSymbolSemiColon:
		return ";"
	case TokenTypeSymbolColon:
		return ":"
	case TokenTypeSymbolQuestion:
		return "?"
	case TokenTypeSymbolComma:
		return ","
	case TokenTypeSymbolPlusPlus:
		return "++"
	case TokenTypeSymbolMinusMinus:
		return "--"
	case TokenTypeSymbolPlusEqual:
		return "+="
	case TokenTypeSymbolDashEqual:
		return "-="
	case TokenTypeSymbolStarEqual:
		return "*="
	case TokenTypeSymbolSlashEqual:
		return "/="
	case TokenTypeSymbolPercentEqual:
		return "%="
	case TokenTypeSymbolPlus:
		return "+"
	case TokenTypeSymbolDash:
		return "-"
	case TokenTypeSymbolStar:
		return "*"
	case TokenTypeSymbolSlash:
		return "/"
	case TokenTypeSymbolPercent:
		return "%"
	case TokenTypeKeywordVar:
		return "var"
	case TokenTypeKeywordLet:
		return "let"
	case TokenTypeKeywordVal:
		return "val"
	case TokenTypeKeywordClass:
		return "class"
	case TokenTypeKeywordNew:
		return "new"
	case TokenTypeKeywordPublic:
		return "public"
	case TokenTypeKeywordPrivate:
		return "private"
	case TokenTypeKeywordStatic:
		return "static"
	case TokenTypeKeywordFinal:
		return "final"
	case TokenTypeKeywordAbstract:
		return "abstract"
	case TokenTypeKeywordOverride:
		return "override"
	case TokenTypeKeywordImport:
		return "import"
	case TokenTypeKeywordFrom:
		return "from"
	case TokenTypeKeywordFunc:
		return "func"
	case TokenTypeKeywordIf:
		return "if"
	case TokenTypeKeywordElse:
		return "else"
	case TokenTypeKeywordElseIf:
		return "elseif"
	case TokenTypeKeywordSwitch:
		return "switch"
	case TokenTypeKeywordCase:
		return "case"
	case TokenTypeKeywordDefault:
		return "default"
	case TokenTypeKeywordForeach:
		return "foreach"
	case TokenTypeKeywordIn:
		return "in"
	case TokenTypeKeywordFor:
		return "for"
	case TokenTypeKeywordWhile:
		return "while"
	case TokenTypeKeywordExport:
		return "export"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}

func newUniqueToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
}
