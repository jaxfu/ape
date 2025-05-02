package internal

import (
	"unicode"
)

func Lex(source string) ([]Token, error) {
	tokens := []Token{}

	// var linenum uint = 0
	workingon := "UNDEFINED"
	curr := ""
	for _, char := range source {
		if workingon {
			switch {
			case unicode.IsSpace(char):
				workingon = false
			default:
				curr += string(char)
			}
		} else {
		}
	}

	return tokens, nil
}

// func Lex(source string) ([]Token, error) {
// 	tokens := []Token{}
//
// 	lines := strings.SplitSeq(source, SYMBOL_FORMATTING_NEWLINE)
// 	var lineNum uint = 0
// 	for line := range lines {
// 		lineNum += 1
// 		line = strings.TrimSpace(line)
// 		split := strings.Fields(line)
//
// 		// NEWLINE
// 		if line == "" {
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().EMPTYLINE, "", lineNum),
// 			)
//
// 			// CONSTRAINT
// 		} else if strings.Contains(line, SYMBOL_DEFINE_CONSTRAINT) {
// 			before, after, found := strings.Cut(line, SYMBOL_DEFINE_CONSTRAINT)
// 			if !found {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
//
// 			key := strings.TrimSpace(before)
// 			if key == "" {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
// 			value := strings.TrimSpace(after)
// 			if value == "" {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().CONSTRAINT_KEY, key, lineNum),
// 			)
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().CONSTRAINT_DEFINITION_SYMBOL, SYMBOL_DEFINE_CONSTRAINT, lineNum),
// 			)
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().CONSTRAINT_VALUE, value, lineNum),
// 			)
//
// 			// DECLARATION
// 		} else if strings.Contains(line, SYMBOL_DEFINE_COMPONENT) {
// 			childOffset := 0
// 			for ; childOffset < len(line); childOffset++ {
// 				if line[childOffset] == '|' {
// 					tokens = append(
// 						tokens,
// 						NewToken(TokenTypes.Types().COMMENT_SYMBOL, SYMBOL_DESCRIBE_UNION, lineNum),
// 					)
// 				} else if line[childOffset] != ' ' &&
// 					line[childOffset] != '\t' {
// 					break
// 				}
// 			}
// 			line = line[childOffset:]
//
// 			before, after, found := strings.Cut(line, SYMBOL_DEFINE_COMPONENT)
// 			if !found {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
// 			name := strings.TrimSpace(before)
// 			value := strings.TrimSpace(after)
//
// 			if name == "" {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
// 			if value == "" {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
//
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().COMPONENT_NAME, name, lineNum),
// 			)
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().COMPONENT_DEFINITION_SYMBOL, SYMBOL_DEFINE_COMPONENT, lineNum),
// 			)
//
// 			// ENUM
// 			if strings.ToLower(value) == SYMBOL_DEFINE_ENUM {
// 				tokens = append(
// 					tokens,
// 					NewToken(TokenTypes.Types().COMPONENT_TYPE_ENUM, value, lineNum),
// 				)
//
// 				// NORMAL
// 			} else {
// 				tokens = append(
// 					tokens,
// 					NewToken(TokenTypes.Types().COMPONENT_TYPE, value, lineNum),
// 				)
// 			}
//
// 			// SOLO ENUM KEY
// 		} else if len(split) == 1 {
// 			key := strings.TrimSpace(split[0])
// 			if key == "" {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().ENUM_KEY, key, lineNum),
// 			)
//
// 			// ENUM KEY w/ NAME
// 		} else if len(split) == 2 {
// 			key := strings.TrimSpace(split[0])
// 			if key == "" {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().ENUM_KEY, key, lineNum),
// 			)
//
// 			value := strings.TrimSpace(split[1])
// 			if value == "" {
// 				return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 			}
// 			tokens = append(
// 				tokens,
// 				NewToken(TokenTypes.Types().ENUM_VALUE, value, lineNum),
// 			)
//
// 			// INVALID
// 		} else {
// 			return nil, fmt.Errorf("invalid syntax line %d", lineNum)
// 		}
// 	}
//
// 	tokens = append(
// 		tokens,
// 		NewToken(TokenTypes.Types().EOF, "EOF", lineNum),
// 	)
//
// 	return tokens, nil
// }
