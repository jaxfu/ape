package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/internal/db"
	"github.com/jaxfu/ape/engine/core/validator"
)

type Store struct {
	Components map[components.ComponentId]components.Component
	Categories map[components.CategoryId]struct{}
	Db         db.Db
}

func NewStore(dbName, initFp string) *Store {
	db, err := db.NewDb(
		dbName,
		initFp,
	)
	if err != nil {
		log.Fatalf("error opening db at %s: %+v\n", dbName, err)
	}

	return &Store{
		Components: map[components.ComponentId]components.Component{},
		Categories: map[components.CategoryId]struct{}{},
		Db:         db,
	}
}

func (s *Store) CreateComponent(comp components.Component) error {
	if err := validator.NewValidator().ValidateComponent(comp); err != nil {
		return fmt.Errorf("error validating: %+v", err)
	}
	if _, ok := s.Components[comp.Metadata().ComponentId]; ok {
		return fmt.Errorf("component with id %s already exists", comp.Metadata().ComponentId)
	}
	s.Components[comp.Metadata().ComponentId] = comp

	marshalled, err := json.Marshal(comp)
	if err != nil {
		return fmt.Errorf("json.Marshal: %+v", err)
	}
	if err := s.Db.InsertComponent(comp.Metadata().ComponentId, marshalled); err != nil {
		return fmt.Errorf("Db.InsertComponent: %+v", err)
	}

	fmt.Println("component stored")
	return nil
}
