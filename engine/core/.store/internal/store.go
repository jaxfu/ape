package internal

import (
	"fmt"

	"github.com/jaxfu/ape/components"
	"github.com/jaxfu/ape/engine/core/store/cache"
	"github.com/jaxfu/ape/engine/core/store/repo"
)

type Store struct {
	Repo           repo.Repository
	ComponentCache cache.ComponentCacheInterface
}

func DefaultStore(r repo.Repository) *Store {
	return &Store{
		Repo:           r,
		ComponentCache: cache.NewComponentCache(),
	}
}

func (s *Store) Add(any) error {
	return nil
}

func (s *Store) Get(name string) (any, error) {
	return nil, nil
}

func (s *Store) Sync(ac components.AllComponents) error {
	// fmt.Printf("%+v\n", ac)
	fmt.Printf("syncing store...\n")

	if len(ac.Props) > 0 {
		fmt.Println("\tProps:")
	}
	for _, v := range ac.Props {
		// wip error
		if err := s.Repo.StoreProp(v); err != nil {
			return fmt.Errorf("Repo.StoreProp: %+v", err)
		}

		if err := s.ComponentCache.AddProp(v); err != nil {
			return fmt.Errorf("Cache.AddProp: %+v", err)
		}

		fmt.Printf("\t\t\u2713 %s\n", v.ComponentMetadata.ComponentId.Display)
	}

	if len(ac.Objects) > 0 {
		fmt.Println("\tObjects:")
	}
	for _, v := range ac.Objects {
		if err := s.Repo.StoreObject(v); err != nil {
			return fmt.Errorf("Repo.StoreObject: %+v", err)
		}

		if err := s.ComponentCache.AddObject(v); err != nil {
			return fmt.Errorf("Cache.AddObject: %+v", err)
		}

		fmt.Printf("\t\t\u2713 %s\n", v.ComponentMetadata.ComponentId.Display)

		for k := range v.Props {
			fmt.Printf("\t\t\t\u2713 %s\n", k)
		}
	}

	fmt.Printf("store synced\n\n")
	return nil
}
