package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	JobTitle           string
	SeniorityLevel     string
	SeniorityLevelRank int
	Country            string
	City               string
	Salary             int
	Currency           string
	RequiredSkills     string
	CompanySize        string
	CompanyDomain      string
}
