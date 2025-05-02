package internal

import (
	"fmt"
	"strings"
	"unicode"
)

// TODO: wip
func Analyze(ast Ast) error {
	for _, node := range ast.RootNodes {
		if err := analyzeDecNode(node); err != nil {
			return fmt.Errorf("error analyzing component '%s': %+v", node.Name, err)
		}
	}

	return nil
}

func analyzeDecNode(node *DeclarationNode) error {
	if err := analyzeStr(
		node.Name,
		true,
		false,
	); err != nil {
		return err
	}

	if err := analyzeStr(
		node.Type,
		false,
		true,
	); err != nil {
		return err
	}

	if strings.TrimSpace(node.Type) == SYMBOL_DEFINE_ENUM {
		if err := analyzeEnums(node.EnumKeys); err != nil {
			return fmt.Errorf("error in enum keys: %+v", err)
		}
	} else {
		if err := analyzeConstraints(node.Constraints); err != nil {
			return fmt.Errorf("error in constraints: %+v", err)
		}
	}

	return nil
}

func analyzeConstraints(cons []*ConstraintNode) error {
	for _, con := range cons {
		fmt.Printf("%+v", con)
		if err := analyzeStr(
			con.Key,
			false,
			false,
		); err != nil {
			return fmt.Errorf("invalid constraint key")
		}

		if err := analyzeStr(
			con.Value,
			true,
			false,
		); err != nil {
			return fmt.Errorf("invalid constraint value")
		}
	}

	return nil
}

func analyzeEnums(enums []EnumKey) error {
	for _, enum := range enums {
		fmt.Printf("%+v", enum)
		if err := analyzeStr(
			enum.Key,
			false,
			false,
		); err != nil {
			return fmt.Errorf("invalid enum key")
		}

		if err := analyzeStr(
			enum.Value,
			true,
			false,
		); err != nil {
			return fmt.Errorf("invalid enum name")
		}
	}

	return nil
}

func analyzeStr(src string, alnum bool, alpha bool) error {
	sani := strings.TrimSpace(src)
	if sani == "" {
		return fmt.Errorf("missing name")
	}

	if alnum {
		if !isAlnum(sani) {
			return fmt.Errorf(
				"'%s' invalid; must be alphanumeric",
				src,
			)
		}
	}

	if alpha {
		if !isAlpha(sani) {
			return fmt.Errorf(
				"'%s' invalid; must be alphabetical",
				src,
			)
		}
	}

	return nil
}

func isAlnum(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return len(s) > 0
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && r != '_' {
			return false
		}
	}
	return len(s) > 0
}
