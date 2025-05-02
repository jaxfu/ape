package shared

const (
	RUNE_CARRIAGE_RETURN   rune = '\r'
	RUNE_LINEFEED          rune = '\n'
	RUNE_TAB               rune = '\t'
	RUNE_SPACE             rune = ' '
	RUNE_UNDERSCORE        rune = '_'
	RUNE_HYPHEN            rune = '-'
	RUNE_GREATER_THAN      rune = '>'
	RUNE_PERIOD            rune = '.'
	RUNE_QUESTION          rune = '?'
	RUNE_COMMENT_SYMBOL    rune = '#'
	RUNE_CONSTRAINT_SYMBOL rune = ':'
)

const (
	SYMBOL_DECLARE_COMPONENT  string = "->"
	SYMBOL_DECLARE_CONSTRAINT string = ":"

	KEYWORD_DECLARE_ENUM  string = "enum"
	KEYWORD_DECLARE_UNION string = "union"
	KEYWORD_DECLARE_ARRAY string = "array"

	SYMBOL_START_COMMENT string = "#"

	SYMBOL_MARK_REFERENCE string = "."
	SYMBOL_MARK_OPTIONAL  string = "?"

	SYMBOL_HYPHEN = "-"
	SYMBOL_GT     = ">"

	SYMBOL_DESCRIBE_UNION string = "|"

	SYMBOL_NEWLINE         string = "\n"
	SYMBOL_CARRIAGE_RETURN string = "\r"
	SPACER_TAB             string = "\t"
	SPACER_SPACE           string = " "

	SYMBOL_SEPARATOR_UNION  string = " | "
	SYMBOL_SEPARATOR_SPACER string = " "
)
