package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Country string `json:"country`
	Savings float64 `json:"savings" gorm:"type:numeric(10,2)"`
}
