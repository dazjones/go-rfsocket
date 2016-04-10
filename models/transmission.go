package models

import (
	"github.com/jinzhu/gorm"
)

type Transmission struct {
	gorm.Model

	Status int
	Rfcode int
	Topic string
}

func (Transmission) TableName() string {
	return "transmission"
}