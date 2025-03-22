package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
)

func (s *Store) GetObjects() ([]components.Object, error) {
	objs, err := s.Repo.GetObjects()
	if err != nil {
		return objs, fmt.Errorf("Repo.GetObjects: %+v", err)
	}

	return objs, nil
}
