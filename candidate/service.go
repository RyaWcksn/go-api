package candidate

type ServiceCandidate interface {
	GetAll() ([]Candidate, error)
	Create(candidate Candidate) (Candidate, error)
}

type service struct {
	repo CandidateRepository
}

func NewCandidateService(repo CandidateRepository) *service {
    return &service{repo}
}

func (s *service) GetAll() ([]Candidate, error) {
    return s.repo.GetAll()
}

func (s *service) Create(candidate Candidate) (Candidate, error) {
    return s.repo.Create(candidate)
}
