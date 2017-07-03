package model

import (
	"time"

	"fmt"

	"github.com/jinzhu/gorm"
)

type (
	Tag struct {
		TagId     int64     `json:"tag_id" gorm:column:tag_id;primary_key`
		TagName   string    `json:"tag_name" gorm:"column:tag_name"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	}

	tagDao struct {
		*gorm.DB
	}

	TagDao interface {
		Dao
		FindByCondition() ([]*Tag, error)
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

/*
条件指定によるタグ一覧取得
*/
func (dao *tagDao) FindByCondition() ([]*Tag, error) {
	var tags []*Tag
	q := dao.table()

	if res := q.Limit(10).Find(&tags); res.Error != nil {
		return nil, fmt.Errorf("failed get tags by condition. {err=%s}", res.Error.Error())
	}

	return tags, nil
}
