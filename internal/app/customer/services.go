package customer

type (
	repoProvider interface {
	}

	Service struct {
		repo repoProvider
	}
)

func NewService(repo repoProvider) *Service {
	return &Service{
		repo: repo,
	}
}
