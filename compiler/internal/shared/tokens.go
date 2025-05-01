package shared

const (
	TOKEN_IDENT       = "IDENT"
	TOKEN_NUMBER      = "NUMBER"
	TOKEN_SYMBOL      = "SYMBOL"
	TOKEN_STRING      = "STRING"
	TOKEN_COMMENT_SYM = "COMMENT_SYM"
	TOKEN_NEWLINE     = "NEWLINE"
	TOKEN_SPACE       = "SPACE"
	TOKEN_TAB         = "TAB"
	TOKEN_EOF         = "EOF"
	TOKEN_UNDEFINED   = "UNDEFINED"
)

// switch tok.Type {
// case shared.TOKEN_IDENT:
// case shared.TOKEN_NUMBER:
// case shared.TOKEN_SYMBOL:
// case shared.TOKEN_STRING:
// case shared.TOKEN_COMMENT_SYM:
// case shared.TOKEN_NEWLINE:
// case shared.TOKEN_SPACE:
// case shared.TOKEN_TAB:
// case shared.TOKEN_EOF:
// case shared.TOKEN_UNDEFINED:
// default:
// }

type TokenType = string

type Token struct {
	Type     TokenType
	Content  string
	Position Position
}

func NewToken(tt TokenType, cnt string, pos Position) Token {
	return Token{
		Type:     tt,
		Content:  cnt,
		Position: pos,
	}
}
