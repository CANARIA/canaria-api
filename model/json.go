package model

type (
	PreRegister struct {
		MailAddress string `json:"mailaddress"`
	}

	AuthRegister struct {
		UserName    string `json:"user_name"`
		MailAddress string `json:"mailaddress"`
		Password    string `json:"password"`
	}
)
