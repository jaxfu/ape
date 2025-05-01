package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

var componentTypes = components.ComponentTypes.Types()

func (v Validator) ValidateComponent(comp components.Component) error {
	meta := comp.Metadata()
	if err := v.validateComponentMetadata(meta); err != nil {
		return fmt.Errorf("error validating component metadata: %+v", err)
	}
	if err := validateComponentType(comp); err != nil {
		return fmt.Errorf("error validating component type for %s: %+v", meta.Name, err)
	}
	if err := v.validateComponentSpecific(comp); err != nil {
		return fmt.Errorf("error validating component %s: %+v", meta.Name, err)
	}

	return nil
}

func (v Validator) validateComponentSpecific(comp components.Component) error {
	switch c := comp.(type) {
	case components.Prop:
		if err := v.validateProp(c); err != nil {
			return fmt.Errorf("error validating prop: %+v", err)
		}
	case components.Object:
		if err := v.validateObject(c); err != nil {
			return fmt.Errorf("error validating object: %+v", err)
		}
	case components.Route:
		if err := v.validateRoute(c); err != nil {
			return fmt.Errorf("error validating route: %+v", err)
		}
	case components.MessageBody:
		if err := v.validateBody(c); err != nil {
			return fmt.Errorf("error validating message body: %+v", err)
		}
	case components.Request:
		if err := v.validateRequest(c); err != nil {
			return fmt.Errorf("error validating request: %+v", err)
		}
	case components.Response:
		if err := v.validateResponse(c); err != nil {
			return fmt.Errorf("error validating response: %+v", err)
		}
	default:
		return fmt.Errorf("unrecognized component type")
	}

	return nil
}

func validateComponentType(comp components.Component) error {
	meta := comp.Metadata()

	switch comp.(type) {
	case components.Prop:
		if meta.ComponentType != componentTypes.PROP {
			return fmt.Errorf("incorrect type for %s: want %s, got %s", meta.ComponentId, componentTypes.PROP, meta.ComponentType)
		}
	case components.Object:
		if meta.ComponentType != componentTypes.OBJECT {
			return fmt.Errorf("incorrect type for %s: want %s, got %s", meta.ComponentId, componentTypes.OBJECT, meta.ComponentType)
		}
	case components.Route:
		if meta.ComponentType != componentTypes.ROUTE {
			return fmt.Errorf("incorrect type for %s: want %s, got %s", meta.ComponentId, componentTypes.ROUTE, meta.ComponentType)
		}
	case components.MessageBody:
		if meta.ComponentType != componentTypes.MESSAGE_BODY {
			return fmt.Errorf("incorrect type for %s: want %s, got %s", meta.ComponentId, componentTypes.MESSAGE_BODY, meta.ComponentType)
		}
	case components.Request:
		if meta.ComponentType != componentTypes.REQUEST {
			return fmt.Errorf("incorrect type for %s: want %s, got %s", meta.ComponentId, componentTypes.REQUEST, meta.ComponentType)
		}
	case components.Response:
		if meta.ComponentType != componentTypes.RESPONSE {
			return fmt.Errorf("incorrect type for %s: want %s, got %s", meta.ComponentId, componentTypes.RESPONSE, meta.ComponentType)
		}
	default:
		return fmt.Errorf("unrecognized component type")
	}

	return nil
}
