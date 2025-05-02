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

const (
	TOKEN_STRING    = "STRING"
	TOKEN_SYMBOL    = "SYMBOL"
	TOKEN_INDENT    = "INDENT"
	TOKEN_NEWLINE   = "NEWLINE"
	TOKEN_EMPTYLINE = "EMPTYLINE"
	TOKEN_SPACE     = "SPACE"
)

var toktypes = Types

type (
	TokenType = string
	Token     struct {
		Type    TokenType
		Value   string
		LineNum uint
	}
)

type TokenTypesInterface struct {
	COMPONENT_NAME              TokenType
	COMPONENT_DEFINITION_SYMBOL TokenType
	COMPONENT_TYPE              TokenType

	CONSTRAINT_NAME              TokenType
	CONSTRAINT_DEFINITION_SYMBOL TokenType
	CONSTRAINT_VALUE             TokenType

	COMPONENT_TYPE_ENUM TokenType
	ENUM_KEY            TokenType
	ENUM_VALUE          TokenType

	COMPONENT_TYPE_UNION TokenType
	UNION_SYMBOL         TokenType

	COMPONENT_TYPE_ARRAY TokenType

	COMMENT_SYMBOL TokenType
	COMMENT_VALUE  TokenType

	REFERENCE_SYMBOL TokenType
	OPTIONAL_SYMBOL  TokenType

	INDENT    TokenType
	EMPTYLINE TokenType
	EOF       TokenType
	UNDEFINED TokenType
}

const (
	COMPONENT_NAME              TokenType = "COMPONENT_NAME"
	COMPONENT_DEFINITION_SYMBOL TokenType = "COMPONENT_DEFINITION_SYMBOL"
	COMPONENT_TYPE              TokenType = "COMPONENT_TYPE"

	CONSTRAINT_NAME              TokenType = "CONSTRAINT_NAME"
	CONSTRAINT_DEFINITION_SYMBOL TokenType = "CONSTRAINT_DEFINITION_SYMBOL"
	CONSTRAINT_VALUE             TokenType = "CONSTRAINT_VALUE"

	COMPONENT_TYPE_ENUM TokenType = "COMPONENT_TYPE_ENUM"
	ENUM_KEY            TokenType = "ENUM_KEY"
	ENUM_VALUE          TokenType = "ENUM_VALUE"

	COMPONENT_TYPE_UNION TokenType = "COMPONENT_TYPE_UNION"
	UNION_SYMBOL         TokenType = "UNION_SYMBOL"

	COMPONENT_TYPE_ARRAY TokenType = "COMPONENT_TYPE_ARRAY"

	COMMENT_SYMBOL TokenType = "COMMENT_SYMBOL"
	COMMENT_VALUE  TokenType = "COMMENT_VALUE"

	REFERENCE_SYMBOL TokenType = "REFERENCE_SYMBOL"
	OPTIONAL_SYMBOL  TokenType = "OPTIONAL_SYMBOL"

	INDENT    TokenType = "INDENT"
	EMPTYLINE TokenType = "EMPTYLINE"
	EOF       TokenType = "EOF"
	UNDEFINED TokenType = "UNDEFINED"
)

var Types = TokenTypesInterface{
	COMPONENT_NAME:               COMPONENT_NAME,
	COMPONENT_DEFINITION_SYMBOL:  COMPONENT_DEFINITION_SYMBOL,
	COMPONENT_TYPE:               COMPONENT_TYPE,
	COMPONENT_TYPE_ENUM:          COMPONENT_TYPE_ENUM,
	ENUM_KEY:                     ENUM_KEY,
	ENUM_VALUE:                   ENUM_VALUE,
	COMPONENT_TYPE_UNION:         COMPONENT_TYPE_UNION,
	UNION_SYMBOL:                 UNION_SYMBOL,
	COMPONENT_TYPE_ARRAY:         COMPONENT_TYPE_ARRAY,
	CONSTRAINT_NAME:              CONSTRAINT_NAME,
	CONSTRAINT_DEFINITION_SYMBOL: CONSTRAINT_DEFINITION_SYMBOL,
	CONSTRAINT_VALUE:             CONSTRAINT_VALUE,
	COMMENT_SYMBOL:               COMMENT_SYMBOL,
	COMMENT_VALUE:                COMMENT_VALUE,
	REFERENCE_SYMBOL:             REFERENCE_SYMBOL,
	OPTIONAL_SYMBOL:              OPTIONAL_SYMBOL,
	INDENT:                       INDENT,
	EMPTYLINE:                    EMPTYLINE,
	EOF:                          EOF,
	UNDEFINED:                    UNDEFINED,
}

var MatchMap = map[string]TokenType{
	string(COMPONENT_NAME):               Types.COMPONENT_NAME,
	string(COMPONENT_DEFINITION_SYMBOL):  Types.COMPONENT_DEFINITION_SYMBOL,
	string(COMPONENT_TYPE):               Types.COMPONENT_TYPE,
	string(COMPONENT_TYPE_ENUM):          Types.COMPONENT_TYPE_ENUM,
	string(ENUM_KEY):                     Types.ENUM_KEY,
	string(ENUM_VALUE):                   Types.ENUM_VALUE,
	string(COMPONENT_TYPE_UNION):         Types.COMPONENT_TYPE_UNION,
	string(UNION_SYMBOL):                 Types.UNION_SYMBOL,
	string(COMPONENT_TYPE_ARRAY):         Types.COMPONENT_TYPE_ARRAY,
	string(CONSTRAINT_NAME):              Types.CONSTRAINT_NAME,
	string(CONSTRAINT_DEFINITION_SYMBOL): Types.CONSTRAINT_DEFINITION_SYMBOL,
	string(CONSTRAINT_VALUE):             Types.CONSTRAINT_VALUE,
	string(COMMENT_SYMBOL):               Types.COMMENT_SYMBOL,
	string(COMMENT_VALUE):                Types.COMMENT_VALUE,
	string(REFERENCE_SYMBOL):             Types.REFERENCE_SYMBOL,
	string(OPTIONAL_SYMBOL):              Types.OPTIONAL_SYMBOL,
	string(INDENT):                       Types.INDENT,
	string(EMPTYLINE):                    Types.EMPTYLINE,
	string(EOF):                          Types.EOF,
	string(UNDEFINED):                    Types.UNDEFINED,
}
