package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/repo/db/sql/generated"
)

func (rp *Repository) StoreProp(prop components.Prop) error {
	// m, _ := json.MarshalIndent(prop, "", "\t")
	// fmt.Println(string(m))

	cat := sql.NullString{}
	if prop.Category != nil && strings.TrimSpace(*prop.Category) != "" {
		cat.String = *prop.Category
		cat.Valid = true
	}
	des := sql.NullString{}
	if prop.Description != nil && strings.TrimSpace(*prop.Description) != "" {
		des.String = *prop.Description
		des.Valid = true
	}

	if err := rp.QueryClient.InsertComponentMetadata(
		context.Background(),
		generated.InsertComponentMetadataParams{
			ComponentID: prop.ComponentId,
			DisplayID:   prop.DisplayId,
			Name:        prop.Name,
			Description: des,
			CategoryID:  cat,
		},
	); err != nil {
		return fmt.Errorf("InsertComponentMetadata: %+v", err)
	}

	parentId := sql.NullString{}
	if prop.ParentId != nil && strings.TrimSpace(*prop.ParentId) != "" {
		parentId.String = *prop.ParentId
		parentId.Valid = true
	}
	if err := rp.QueryClient.InsertProp(
		context.Background(),
		generated.InsertPropParams{
			ComponentID: prop.ComponentId,
			PropType:    prop.PropType,
			ParentID:    parentId,
		},
	); err != nil {
		return fmt.Errorf("InsertPropRelationship: %+v", err)
	}

	// opts
	if err := rp.storeOpts(prop, prop.ComponentId); err != nil {
		return fmt.Errorf("Repo.storeOpts: %+v", err)
	}

	// check if ref'd
	refees, err := rp.QueryClient.GetRefsByTargetId(
		context.Background(),
		prop.ComponentId,
	)
	if err != nil {
		return fmt.Errorf("Repo.GetRefsByTargetId: %+v", err)
	}

	for _, v := range refees {
		if err := rp.QueryClient.UpdateRef(
			context.Background(),
			generated.UpdateRefParams{
				LinkedTargetID: sql.NullString{
					String: prop.ComponentId,
					Valid:  true,
				},
				ComponentID: v.ComponentID,
			},
		); err != nil {
			fmt.Printf("error updating Ref '%s to '%s': %+v\n", v.ComponentID, prop.DisplayId, err)
		}
	}

	return nil
}

func (rp *Repository) GetProps() ([]components.Prop, error) {
	props := []components.Prop{}

	fetched, err := rp.QueryClient.GetProps(context.Background())
	if err != nil {
		return props, fmt.Errorf("QueryClient.GetProps: %+v", err)
	}

	for _, v := range fetched {
		cat := new(string)
		if v.CategoryID.String != "" {
			cat = &v.CategoryID.String
		} else {
			cat = nil
		}

		des := new(string)
		if v.Description.String != "" {
			des = &v.Description.String
		} else {
			des = nil
		}

		pId := new(string)
		if v.ParentID.String != "" {
			pId = &v.ParentID.String
		}

		prop := components.Prop{
			ComponentMetadata: components.ComponentMetadata{
				ComponentId:   v.ComponentID,
				ComponentType: components.COMPONENT_TYPE_PROP,
				ParentId:      pId,
				Name:          v.Name,
				Category:      cat,
				Description:   des,
			},
			PropType: v.PropType.String,
			Opts:     components.Opts{},
		}

		props = append(props, prop)
	}

	return props, nil
}

func (rp *Repository) getPropsByObject(id components.ComponentId) (*components.PropsMap, error) {
	pm := &components.PropsMap{}

	parentId := sql.NullString{
		String: strings.ToLower(id),
		Valid:  true,
	}

	props, err := rp.QueryClient.GetPropsByParentId(
		context.Background(),
		parentId,
	)
	if err != nil {
		return nil, err
	}

	for _, v := range props {
		cat := new(string)
		if v.CategoryID.String != "" {
			*cat = v.CategoryID.String
		} else {
			cat = nil
		}
		des := new(string)
		if v.Description.String != "" {
			*des = v.Description.String
		} else {
			des = nil
		}

		// opts
		opts, err := rp.getOpts(v.ComponentID, v.PropType)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		prop := components.Prop{
			ComponentMetadata: components.ComponentMetadata{
				Name:          v.Name,
				ComponentId:   v.ComponentID,
				ComponentType: components.COMPONENT_TYPE_PROP,
				Category:      cat,
				Description:   des,
			},
			PropType: v.PropType,
			Opts:     opts,
		}

		(*pm)[v.Name] = prop
	}

	return pm, nil
}
