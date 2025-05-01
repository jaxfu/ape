package assembler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jaxfu/ape/compiler/internal/shared"
	"github.com/jaxfu/ape/components"
)

var keywordToComponentTypeMap = map[string]components.ComponentType{
	shared.KEYWORD_OBJECT: components.COMPONENT_TYPE_OBJECT,
	shared.KEYWORD_ENUM:   components.COMPONENT_TYPE_ENUM,
	shared.KEYWORD_ARRAY:  components.COMPONENT_TYPE_ARRAY,
	shared.KEYWORD_STRING: components.COMPONENT_TYPE_STRING,
	shared.KEYWORD_BLOB:   components.COMPONENT_TYPE_BLOB,
	shared.KEYWORD_INT:    components.COMPONENT_TYPE_INT,
	shared.KEYWORD_UINT:   components.COMPONENT_TYPE_UINT,
	shared.KEYWORD_FLOAT:  components.COMPONENT_TYPE_FLOAT,
	shared.KEYWORD_BOOL:   components.COMPONENT_TYPE_BOOL,
}

func processNode(node shared.Node, ctx AssemblyCtx) AssemblyCtx {
	switch n := node.(type) {
	case shared.NodeComponent:
		ctx = processNodeComponent(n, ctx)
	case shared.NodeEnumMember:
		ctx = processNodeEnumMember(n, ctx)
	case shared.NodeTrait:
		ctx = processNodeTrait(n, ctx)
	default:
		ctx.Error = shared.NewSyntaxError(node.Meta().Position, "unrecognized node type")
	}

	return ctx
}

func classifyNodeComponent(
	node shared.NodeComponent,
) (components.ComponentType, error) {
	if node.IsReference {
		return components.COMPONENT_TYPE_REFERENCE, nil
	}

	ctype, ok := keywordToComponentTypeMap[strings.ToLower(node.Type)]
	if ok {
		return ctype, nil
	} else {
		return components.COMPONENT_TYPE_UNDEFINED, shared.NewSyntaxError(node.Meta().Position, fmt.Sprintf("invalid component type '%s'\n", node.Type))
	}
}

func processNodeComponent(
	node shared.NodeComponent,
	ctx AssemblyCtx,
) AssemblyCtx {
	ctype, err := classifyNodeComponent(node)
	if err != nil {
		ctx.Error = err
		return ctx
	}

	comp := components.NewComponent(ctype, "test_id", findParent(ctx))
	ctx.Comps[comp.Meta().ComponentId] = &comp
	ctx.Stack.Push(&comp)

	return ctx
}

func processNodeTrait(
	node shared.NodeTrait,
	ctx AssemblyCtx,
) AssemblyCtx {
	parent := findParent(ctx)
	if parent == nil {
		ctx.Error = shared.NewSyntaxError(node.Meta().Position, "orphaned trait")
		return ctx
	}

	meta := components.TraitMetadata{
		Key: node.Name,
		// TODO: wip fix, pass type from parser
		Type: components.TRAIT_TYPE_UINT,
		Raw:  node.Value,
	}

	_, err := newTrait(meta)
	if err != nil {
		ctx.Error = err
		return ctx
	}

	return ctx
}

func processNodeEnumMember(
	node shared.NodeEnumMember,
	ctx AssemblyCtx,
) AssemblyCtx {
	return ctx
}

func findParent(ctx AssemblyCtx) *components.Component {
	curr, ln := ctx.Stack.Curr()
	if ln > 0 {
		return curr
	} else {
		return nil
	}
}

func newTrait(meta components.TraitMetadata) (components.Trait, error) {
	switch meta.Type {
	case components.TRAIT_TYPE_STRING:
		return components.TraitString{
			TraitMetadata: meta,
			Value:         meta.Raw,
		}, nil
	case components.TRAIT_TYPE_INT:
		val, err := castInt(meta.Raw)
		if err != nil {
			return nil, fmt.Errorf("could not cast trait '%s' to %s", meta.Key, meta.Type)
		}

		return components.TraitInt{
			TraitMetadata: meta,
			Value:         val,
		}, nil
	case components.TRAIT_TYPE_UINT:
		val, err := castUint(meta.Raw)
		if err != nil {
			return nil, fmt.Errorf("could not cast trait '%s' to %s", meta.Key, meta.Type)
		}

		return components.TraitUint{
			TraitMetadata: meta,
			Value:         val,
		}, nil

	case components.TRAIT_TYPE_FLOAT:
		val, err := castFloat(meta.Raw)
		if err != nil {
			return nil, fmt.Errorf("could not cast trait '%s' to %s", meta.Key, meta.Type)
		}

		return components.TraitFloat{
			TraitMetadata: meta,
			Value:         val,
		}, nil
	case components.TRAIT_TYPE_BOOL:
		val, err := castBool(meta.Raw)
		if err != nil {
			return nil, fmt.Errorf("could not cast trait '%s' to %s", meta.Key, meta.Type)
		}

		return components.TraitBool{
			TraitMetadata: meta,
			Value:         val,
		}, nil
	default:
		return nil, fmt.Errorf("invalid trait value '%s' for traint '%s'\n", meta.Raw, meta.Key)
	}
}

func castInt(raw string) (int, error) {
	return strconv.Atoi(raw)
}

func castUint(raw string) (uint, error) {
	cast, err := strconv.ParseUint(raw, 10, 0)
	return uint(cast), err
}

func castFloat(raw string) (float64, error) {
	return strconv.ParseFloat(raw, 0)
}

func castBool(raw string) (bool, error) {
	return strconv.ParseBool(raw)
}
