package shared

import "fmt"

type SyntaxError = error

// position and optional message (only uses first string arg)
func NewSyntaxError(pos Position, msg ...string) SyntaxError {
	if len(msg) == 0 {
		return fmt.Errorf(
			"line %d, col %d: syntax error",
			pos.Line,
			pos.Col,
		)
	}

	return fmt.Errorf(
		"line %d, col %d: syntax error: %s",
		pos.Line,
		pos.Col,
		msg[0],
	)
}

func processingError(line, col uint, msg string) error {
	return fmt.Errorf(
		"line %d, col %d: syntax error: %s",
		line,
		col,
		msg,
	)
}
