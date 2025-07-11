package shared

const (
	RUNE_CARRIAGE_RETURN        rune = '\r'
	RUNE_LINEFEED               rune = '\n'
	RUNE_TAB                    rune = '\t'
	RUNE_SPACE                  rune = ' '
	RUNE_UNDERSCORE             rune = '_'
	RUNE_HYPHEN                 rune = '-'
	RUNE_GREATER_THAN           rune = '>'
	RUNE_PERIOD                 rune = '.'
	RUNE_QUESTION               rune = '?'
	RUNE_POUND                  rune = '#'
	RUNE_COLON                  rune = ':'
	RUNE_OPEN_STRAIGHT_BRACKET  rune = '['
	RUNE_CLOSE_STRAIGHT_BRACKET rune = ']'

	RUNE_SYM_COMMENT        rune = RUNE_POUND
	RUNE_SYM_TYPEDEF        rune = RUNE_COLON
	RUNE_SYM_REFERENCE      rune = RUNE_PERIOD
	RUNE_SYM_OPTIONAL       rune = RUNE_QUESTION
	RUNE_SYM_START_ARRAY    rune = RUNE_OPEN_STRAIGHT_BRACKET
	RUNE_SYM_END_ARRAY      rune = RUNE_CLOSE_STRAIGHT_BRACKET
	RUNE_SYM_ANON_COMPONENT rune = RUNE_UNDERSCORE
)

const (
	KEYWORD_TYPE_OBJECT string = "object"
	KEYWORD_TYPE_ENUM   string = "enum"
	KEYWORD_TYPE_STRING string = "string"
	KEYWORD_TYPE_BLOB   string = "blob"
	KEYWORD_TYPE_INT    string = "int"
	KEYWORD_TYPE_UINT   string = "uint"
	KEYWORD_TYPE_FLOAT  string = "float"
	KEYWORD_TYPE_BOOL   string = "bool"

	SYMBOL_REFERENCE      string = "."
	SYMBOL_OPTIONAL       string = "?"
	SYMBOL_TYPEDEF        string = "::"
	SYMBOL_COMMENT        string = "#"
	SYMBOL_START_ARRAY    string = "["
	SYMBOL_END_ARRAY      string = "]"
	SYMBOL_ANON_COMPONENT string = "_"

	SYMBOL_NEWLINE         string = "\n"
	SYMBOL_CARRIAGE_RETURN string = "\r"
	SYMBOL_TAB             string = "\t"
	SYMBOL_SPACE           string = " "
)
