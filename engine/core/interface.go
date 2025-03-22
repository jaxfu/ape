package core

// type Core struct {
// 	Store  store.StoreInterface
// 	Server api.ApiInterface
// }
//
// func InitCore(db *sql.DB, url string, port uint, clientDirFp string) (*Core, error) {
// 	repo := repo.NewRepository(db)
// 	store := store.NewStore(repo)
// 	server, err := api.NewServer(url, port, clientDirFp, store)
// 	if err != nil {
// 		return nil, fmt.Errorf("generators.NewGenerators: %+v", err)
// 	}
//
// 	return &Core{
// 		Store:  store,
// 		Server: server,
// 	}, nil
// }
