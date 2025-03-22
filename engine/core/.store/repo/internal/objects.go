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

func (rp *Repository) StoreObject(obj components.Object) error {
	cat := sql.NullString{}
	if obj.Category != nil {
		cat.String = *obj.Category
		cat.Valid = true
	}
	des := sql.NullString{}
	if obj.Description != nil {
		des.String = *obj.Description
		des.Valid = true
	}

	// meta
	if err := rp.QueryClient.InsertComponentMetadata(
		context.Background(),
		generated.InsertComponentMetadataParams{
			ComponentID: obj.ComponentId,
			DisplayID:   obj.DisplayId,
			Name:        obj.Name,
			Description: des,
			CategoryID:  cat,
		},
	); err != nil {
		return fmt.Errorf("InsertComponentMetadata: %+v", err)
	}

	// object
	if err := rp.QueryClient.InsertObject(context.Background(), obj.ComponentId); err != nil {
		return fmt.Errorf("InsertObject: %+v", err)
	}

	// props
	for _, v := range obj.Props {
		if err := rp.StoreProp(v); err != nil {
			return fmt.Errorf("Repo.StoreProp: %+v", err)
		}
	}

	return nil
}

func (rp *Repository) GetObjects() ([]components.Object, error) {
	objs := []components.Object{}

	dbObjs, err := rp.QueryClient.GetObjects(context.Background())
	if err != nil {
		return objs, err
	}

	for _, o := range dbObjs {
		cat := new(string)
		if o.CategoryID.String == "" {
			cat = nil
		} else {
			cat = &o.CategoryID.String
		}
		des := new(string)
		if o.Description.String == "" {
			des = nil
		} else {
			des = &o.Description.String
		}

		obj := components.Object{
			ComponentMetadata: components.ComponentMetadata{
				Name:          o.Name,
				ComponentId:   o.ComponentID,
				DisplayId:     o.DisplayID,
				ComponentType: components.COMPONENT_TYPE_OBJECT,
				Category:      cat,
				Description:   des,
			},
		}

		// props
		props, err := rp.getPropsByObject(strings.ToLower(o.ComponentID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return objs, err
		}
		obj.Props = *props

		objs = append(objs, obj)
	}

	// fmt.Printf("%+v\n", objs)
	return objs, nil
}
