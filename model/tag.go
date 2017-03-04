package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Tag struct {
		TagId     int64     `gorm:column:tag_id;primary_key`
		TagName   string    `gorm:"column:tag_name"`
		CreatedAt time.Time `gorm:"column:created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
	}
	tagDao struct {
		*gorm.DB
	}

	TagDao interface {
		Dao
		// ProfileImpl(*AuthRegister, *Account) *Profile
		// Create(*Profile) error
	}
)

func TagDaoFactory(db *gorm.DB) TagDao {
	return &tagDao{
		DB: db,
	}
}

//--------------------------------------------
// Implementations for Dao interface
//--------------------------------------------

func (dao *tagDao) table() *gorm.DB {
	return dao.Table("tags")
}
