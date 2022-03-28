package candidate

type CandidateRepository interface {
   GetAll() ([]Candidate, error) 
   Create(candidate Candidate) error
}
