package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type (
	PreRegister struct {
		MailAddress string `json:"mailaddress" validate:"required,email"`
	}

	CheckToken struct {
		UrlToken string `json:"url_token" validate:"required"`
	}

	AuthRegister struct {
		UrlToken    string `json:"url_token" validate:"required"`
		MailAddress string `json:"mailaddress" validate:"required,email"`
		UserName    string `json:"user_name" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}

	UserInfo struct {
		UserID      int64  `json:"user_id"`
		UserName    string `json:"user_name"`
		DisplayName string `json:"display_name"`
		MailAddress string `json:"mailaddress"`
		Avatar      string `json:"avatar"`
		Roll        int16  `json:"roll"`
		jwt.StandardClaims
	}
)
