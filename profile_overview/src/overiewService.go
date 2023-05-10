package profile_overview

type overviewService struct {
	repository OverviewRepository
}

type OverviewService interface {
}

func NewOverviewService(repository OverviewRepository) OverviewService {
	return &overviewService{
		repository: repository,
	}
}
