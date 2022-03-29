package candidate

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name        string
	Email       string
	Phone       string
	CV          string
	CoverLetter string
	JobID       uint
	Details     string
}
