package candidate

import (
	"github.com/jinzhu/gorm"
)


type Candidate struct {
    gorm.Model
	Name        string
	Email       string
	Phone       string
	CV          string
	CoverLetter string
	JobID       uint
}
