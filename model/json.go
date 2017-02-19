package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type (
	// PreRegister is ユーザ仮登録用
	PreRegister struct {
		MailAddress string `json:"mailaddress" validate:"required,email"`
	}

	// CheckToken is トークンチェック用
	CheckToken struct {
		UrlToken string `json:"url_token" validate:"required"`
	}

	// AuthRegister is ユーザ本登録用
	AuthRegister struct {
		UrlToken    string `json:"url_token" validate:"required"`
		MailAddress string `json:"mailaddress" validate:"required,email"`
		UserName    string `json:"user_name" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}

	// LoginClaim is ログイン要求
	LoginClaim struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}

	// UserInfo is 認証済みユーザ情報
	UserInfo struct {
		UserID      int64  `json:"user_id" db:"user_id"`
		UserName    string `json:"user_name" db:"user_name"`
		DisplayName string `json:"display_name" db:"display_name"`
		MailAddress string `json:"mailaddress" db:"mailaddress"`
		Avatar      string `json:"avatar" db:"avatar"`
		Roll        int16  `json:"roll" db:"roll"`
		jwt.StandardClaims
	}
)
