package repository

type gophermartRepo struct {
	storage IStorage
}

func newGophermartRepo(storage IStorage) *gophermartRepo {
	return &gophermartRepo{storage: storage}
}
