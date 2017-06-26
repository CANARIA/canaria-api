package service

import (
	"fmt"

	"github.com/CANARIA/canaria-api/core/model"
	"github.com/jinzhu/gorm"
)

/*
タグ一覧の取得
*/
func FindTags(tx *gorm.DB) ([]*model.Tag, error) {

	tagDao := model.TagDaoFactory(tx)
	tags, err := tagDao.FindByCondition()

	if err != nil {
		return nil, fmt.Errorf("failed get tags by condition. err={%s}", err.Error())
	}

	return tags, nil
}
