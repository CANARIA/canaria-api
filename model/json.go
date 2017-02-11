package model

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
)
