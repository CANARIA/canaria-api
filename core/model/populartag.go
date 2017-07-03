package model

import (
	"time"

	"fmt"

	"github.com/jinzhu/gorm"
)

type (
	PopularTag struct {
		TagId     int64     `json:"tag_id" gorm:column:tag_id`
		TagName   string    `json:"tag_name" gorm:"column:tag_name"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	}
	popularTagDao struct {
		*gorm.DB
	}

	PopularTagDao interface {
		Dao
		// ProfileImpl(*AuthRegister, *Account) *Profile
		FindAll() ([]*PopularTag, error)
		// Create(*Profile) error
	}
)

func PopularTagDaoFactory(db *gorm.DB) PopularTagDao {
	return &popularTagDao{
		DB: db,
	}
}

//--------------------------------------------
// Implementations for Dao interface
//--------------------------------------------

func (dao *popularTagDao) table() *gorm.DB {
	return dao.Table("popular_tags")
}

/*
すべての人気のタグを取得する(10件)
*/
func (dao *popularTagDao) FindAll() ([]*PopularTag, error) {
	var popularTags []*PopularTag

	if res := dao.table().Find(&popularTags).Limit(10); res.Error != nil {
		return nil, fmt.Errorf("failed get all popular tags {err=%s}", res.Error.Error())
	}

	return popularTags, nil
}
