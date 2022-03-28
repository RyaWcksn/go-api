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

func (r *repository) GetAll() ([]Candidate, error) {
    var candidates []Candidate
    err := r.db.Find(&candidates).Error
    return candidates, err
}

func (r *repository) Create(candidate Candidate) error {
    return r.db.Create(&candidate).Error
}
