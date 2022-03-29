package company

import "gorm.io/gorm"

type Company struct {
    gorm.Model
    Name string
    Address string
    CompanyType int
}
