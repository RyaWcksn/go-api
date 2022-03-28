package candidate

import "gorm.io/gorm"


type CandidateRepository interface {
   GetAll() ([]Candidate, error) 
   Create(candidate Candidate) error
}

type repository struct {
    db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) *repository {
    return &repository{db}
}


