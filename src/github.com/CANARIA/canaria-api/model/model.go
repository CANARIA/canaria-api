package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Dao interface {
		table() *gorm.DB
	}
)
