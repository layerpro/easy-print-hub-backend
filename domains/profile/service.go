package profile

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetProfile(id int) (*Profile, error) {
	profile, err := s.repo.GetProfile(id)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
