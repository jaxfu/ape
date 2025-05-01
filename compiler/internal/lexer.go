package internal

import (
	"fmt"
	"strings"
)

const (
	CONSTRAINT_SYMBOL  string = ":"
	DECLARATION_SYMBOL string = "->"
	CHILD_PREFIX       string = "|"
	NEWLINE            string = "\n"
)

func Lex(source string) ([]Token, error) {
	tokens := []Token{}

	lines := strings.SplitSeq(source, NEWLINE)
	var lineNum uint = 0
	for line := range lines {
		lineNum += 1
		line = strings.TrimSpace(line)

		if line == "" { // newline
			tokens = append(
				tokens,
				NewToken(TokenTypes.Types().EMPTYLINE, "", lineNum),
			)
		} else if strings.Contains(line, CONSTRAINT_SYMBOL) { // constraint
			before, after, found := strings.Cut(line, CONSTRAINT_SYMBOL)
			if !found {
				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
			}

			key := strings.TrimSpace(before)
			value := strings.TrimSpace(after)
			tokens = append(
				tokens,
				NewToken(TokenTypes.Types().CONSTRAINT_KEY, key, lineNum),
			)
			tokens = append(
				tokens,
				NewToken(TokenTypes.Types().CONSTRAINT_ASSIGNMENT, CONSTRAINT_SYMBOL, lineNum),
			)
			tokens = append(
				tokens,
				NewToken(TokenTypes.Types().CONSTRAINT_VALUE, value, lineNum),
			)

		} else if strings.Contains(line, DECLARATION_SYMBOL) { // declaration
			childOffset := 0
			for ; childOffset < len(line); childOffset++ {
				if line[childOffset] != '|' {
					break
				}

				tokens = append(
					tokens,
					NewToken(TokenTypes.Types().CHILD_PREFIX, CHILD_PREFIX, lineNum),
				)
			}
			line = line[childOffset:]

			before, after, found := strings.Cut(line, DECLARATION_SYMBOL)
			if !found {
				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
			}
			before = strings.TrimSpace(before)
			after = strings.TrimSpace(after)
			tokens = append(
				tokens,
				NewToken(TokenTypes.Types().DECLARATION_NAME, before, lineNum),
			)
			tokens = append(
				tokens,
				NewToken(TokenTypes.Types().DECLARATION_ASSIGNMENT, DECLARATION_SYMBOL, lineNum),
			)
			tokens = append(
				tokens,
				NewToken(TokenTypes.Types().DECLARATION_VALUE, after, lineNum),
			)

		} else { // invalid
			return nil, fmt.Errorf("invalid syntax line %d", lineNum)
		}
	}

	tokens = append(
		tokens,
		NewToken(TokenTypes.Types().EOF, "", lineNum),
	)

	return tokens, nil
}
