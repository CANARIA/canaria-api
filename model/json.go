package model

type (
	PreRegister struct {
		MailAddress string `json:"mailaddress"`
	}

	CheckToken struct {
		UrlToken string `json:"url_token"`
	}

	AuthRegister struct {
		UrlToken    string `json:"url_token"`
		UserName    string `json:"user_name"`
		MailAddress string `json:"mailaddress"`
		Password    string `json:"password"`
	}
)
