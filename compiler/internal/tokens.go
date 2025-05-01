package internal

import "github.com/jaxfu/ape/pkg/enum"

func NewToken(toktype TokenType, val string, lineNum uint) Token {
	return Token{
		Type:    toktype,
		Value:   val,
		LineNum: lineNum,
	}
}

var TokenTypes = enum.Enum[TokenType, TokenTypesInterface]{
	TypeList: Types,
	MatchMap: MatchMap,
}

type (
	TokenType string
	Token     struct {
		Type    TokenType
		Value   string
		LineNum uint
	}
)

type TokenTypesInterface struct {
	DECLARATION_NAME       TokenType
	DECLARATION_VALUE      TokenType
	DECLARATION_ASSIGNMENT TokenType
	CONSTRAINT_KEY         TokenType
	CONSTRAINT_ASSIGNMENT  TokenType
	CONSTRAINT_VALUE       TokenType
	CHILD_PREFIX           TokenType
	EMPTYLINE              TokenType
	EOF                    TokenType
	UNDEFINED              TokenType
}

var Types = TokenTypesInterface{
	DECLARATION_NAME:       "DECLARATION_NAME",
	DECLARATION_VALUE:      "DECLARATION_VALUE",
	DECLARATION_ASSIGNMENT: "DECLARATION_ASS",
	CONSTRAINT_KEY:         "CONSTRAINT_KEY",
	CONSTRAINT_ASSIGNMENT:  "CONSTRAINT_ASS",
	CONSTRAINT_VALUE:       "CONSTRAINT_VALUE",
	CHILD_PREFIX:           "CHILD_PREFIX",
	EMPTYLINE:              "EMPTYLINE",
	EOF:                    "EOF",
	UNDEFINED:              "UNDEFINED",
}

var MatchMap = map[string]TokenType{
	"DECLARATION_NAME":  Types.DECLARATION_NAME,
	"DECLARATION_VALUE": Types.DECLARATION_VALUE,
	"DECLARATION_ASS":   Types.DECLARATION_ASSIGNMENT,
	"CONSTRAINT_KEY":    Types.CONSTRAINT_KEY,
	"CONSTRAINT_ASS":    Types.CONSTRAINT_ASSIGNMENT,
	"CONSTRAINT_VALUE":  Types.CONSTRAINT_VALUE,
	"CHILD_PREFIX":      Types.CHILD_PREFIX,
	"EMPTYLINE":         Types.EMPTYLINE,
	"EOF":               Types.EOF,
	"UNDEFINED":         Types.UNDEFINED,
}
