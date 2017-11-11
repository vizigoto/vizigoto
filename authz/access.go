package auth

type AccessService interface {
	Get(id string) (g *Group, err error)
}

type accessService struct {
	repo Repository
}

func NewAccessService(repo Repository) AccessService {
	return &accessService{repo}
}

func (s *accessService) Get(id string) (g *Group, err error) {
	return s.repo.Get(id)
}
