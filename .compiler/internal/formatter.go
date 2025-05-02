package internal

import (
	"fmt"
	"strings"
)

func format(ast Ast, indent bool) ([]byte, error) {
	if ast.RootNodes == nil {
		return nil, fmt.Errorf("null ast")
	}

	lines := []string{}

	for _, node := range ast.RootNodes {
		formatted := formatDeclaration(node, indent)
		lines = append(lines, formatted...)
	}

	return []byte(strings.Join(lines, SYMBOL_FORMATTING_NEWLINE)), nil
}

func formatDeclaration(node *DeclarationNode, indent bool) []string {
	lines := []string{}

	// dec header
	header := formatDeclarationHeader(node)
	lines = append(lines, header)

	cons := []string{}
	if strings.ToLower(node.Type) == SYMBOL_DEFINE_ENUM {
		// if enum
		cons = append(cons, formatEnumKeys(node, indent)...)
	} else {
		// else normal constraint
		cons = append(cons, formatConstraints(node, indent)...)
	}
	lines = append(lines, cons...)
	lines = append(lines, "")

	// children
	for _, c := range node.Children {
		lines = append(lines, formatDeclaration(c, indent)...)
	}

	return lines
}

func formatDeclarationHeader(node *DeclarationNode) string {
	if node.Depth <= 0 {
		return strings.Join([]string{node.Name, SYMBOL_DEFINE_COMPONENT, node.Type}, SYMBOL_SEPARATOR_SPACER)
	}

	prefixes := []string{}
	for range node.Depth {
		prefixes = append(prefixes, SYMBOL_DESCRIBE_UNION)
	}
	prefix := strings.Join(prefixes, SYMBOL_SEPARATOR_UNION)

	out := strings.Join(
		[]string{prefix, node.Name, SYMBOL_DEFINE_COMPONENT, node.Type},
		SYMBOL_SEPARATOR_SPACER,
	)
	return out
}

func formatConstraints(node *DeclarationNode, indent bool) []string {
	indstr := ""
	if indent {
		indstr += strings.Repeat(
			SYMBOL_FORMATTING_INDENT,
			node.Depth+1,
		)
	}

	maxlen := 0
	for _, cons := range node.Constraints {
		keylen := len(cons.Key)
		if keylen > maxlen {
			maxlen = keylen
		}
	}

	lines := []string{}
	for _, cons := range node.Constraints {
		line := indstr
		line += cons.Key
		line += ": "
		line += cons.Value
		line = leftJustify(line, maxlen)
		lines = append(lines, line)
	}

	return lines
}

func formatEnumKeys(node *DeclarationNode, indent bool) []string {
	indstr := ""
	if indent {
		indstr += strings.Repeat(
			SYMBOL_FORMATTING_INDENT,
			node.Depth+1,
		)
	}

	maxlen := 0
	for _, cons := range node.EnumKeys {
		keylen := len(cons.Key)
		if keylen > maxlen {
			maxlen = keylen
		}
	}

	lines := []string{}
	for _, enum := range node.EnumKeys {
		line := indstr
		// if has named value
		if enum.Value != "" {
			line += leftJustify(enum.Key, maxlen)
			line += " "
			line += enum.Value
		} else {
			line += enum.Key
		}
		lines = append(lines, line)
	}

	return lines
}

func leftJustify(src string, totlen int) string {
	padding := totlen - len(src)
	if padding > 0 {
		return src + strings.Repeat(" ", padding)
	}
	return src
}
