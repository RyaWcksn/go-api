package candidate

import "errors"

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
    if candidate.Name == "" {
        return Candidate{}, errors.New("Candidate name is required")
    }
    if candidate.Email == "" {
        return Candidate{}, errors.New("Candidate email is required")
    }
    if candidate.Phone == "" {
        return Candidate{}, errors.New("Candidate phone is required")
    }
    return s.repo.Create(candidate)
}
