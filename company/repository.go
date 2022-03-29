package company

import "gorm.io/gorm"

type CompanyRepository interface {
    GetAll() ([]Company, error)
    Create(company Company) (Company, error)
}

type companyRepository struct {
    db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *companyRepository {
    return &companyRepository{db}
}

func (r *companyRepository) GetAll() ([]Company, error) {
    var companies []Company
    err := r.db.Find(&companies).Error
    return companies, err
}

func (r *companyRepository) Create(company Company) (Company, error) {
    err := r.db.Create(&company).Error
    return company, err
}
