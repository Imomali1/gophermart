package repository

type IStorage interface {
}

type IGophermartRepo interface {
}

type Repository struct {
	Gophermart IGophermartRepo
}

func New(storage IStorage) Repository {
	return Repository{
		Gophermart: newGophermartRepo(storage),
	}
}
