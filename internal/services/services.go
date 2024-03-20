package services

type IGophermartService interface {
}

type ServiceManager struct {
	Gophermart IGophermartService
}

func New(repo IRepository) ServiceManager {
	return ServiceManager{
		Gophermart: newGophermartService(repo),
	}
}
