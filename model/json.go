package model

type (
	AuthRegister struct {
		UserName    string `json:"user_name"`
		MailAddress string `json:"mailaddress"`
		Password    string `json:"password"`
	}
)
