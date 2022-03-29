package candidate

import "gorm.io/gorm"


type CandidateRepository interface {
   GetAll() ([]Candidate, error) 
   Create(candidate Candidate) (Candidate, error)
}

type repository struct {
    db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) *repository {
    return &repository{db}
}

func (r *repository) GetAll() ([]Candidate, error) {
    var candidates []Candidate
    err := r.db.Find(&candidates).Error
    return candidates, err
}

func (r *repository) Create(candidate Candidate) (Candidate, error) {
    err := r.db.Create(&candidate).Error
    return candidate, err
}
