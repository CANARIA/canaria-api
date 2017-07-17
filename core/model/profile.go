package model

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/jinzhu/gorm"
)

type (
	Profile struct {
		DisplayName string         `gorm:"column:display_name"`
		Bio         sql.NullString `gorm:"column:bio"`
		Avatar      sql.NullString `gorm:"column:avatar"`
		Url         sql.NullString `gorm:"column:url"`
		UserId      int64          `gorm:"column:user_id"`
		CreatedAt   time.Time      `gorm:"column:created_at"`
		UpdatedAt   time.Time      `gorm:"column:updated_at"`
		IsDeleted   bool           `gorm:"column:is_deleted"`
	}
	profileDao struct {
		*gorm.DB
	}

	ProfileDao interface {
		Dao
		ProfileImpl(*AuthRegister, *Account) *Profile
		Create(*Profile) error
	}
)

func ProfileDaoFactory(db *gorm.DB) ProfileDao {
	return &profileDao{
		DB: db,
	}
}

//--------------------------------------------
// Implementations for Dao interface
//--------------------------------------------

func (dao *profileDao) table() *gorm.DB {
	return dao.Table("profiles")
}

//--------------------------------------------
// Implementations for Model interface
//--------------------------------------------

func (dao *profileDao) ProfileImpl(authRegister *AuthRegister, account *Account) *Profile {
	return &Profile{
		DisplayName: authRegister.UserName,
		//UserId:      account.UserId,
		CreatedAt:   time.Now(),
	}
}

func (dao *profileDao) Create(profile *Profile) error {

	if res := dao.table().Save(profile); res.Error != nil {
		return fmt.Errorf("failed account create{%v}", profile)
	}

	return nil
}
