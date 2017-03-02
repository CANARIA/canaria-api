package api

import (
	"fmt"

	"net/http"

	"github.com/CANARIA/canaria-api/middleware"
	"github.com/CANARIA/canaria-api/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

/*
TODO:
ミドルウェアで共通的に処理をする
parseしてユーザー情報をctxに突っ込むところまでを共通化する
*/
func CreateTokenString() string {
	// User情報をtokenに込める(DBから取得したユーザー情報を詰める)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &model.UserInfo{
		UserID:      int64(1),
		UserName:    "system",
		DisplayName: "システム",
		MailAddress: "admin@canaria.io",
		Roll:        1,
	})
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenstring
}

func PopularTags() echo.HandlerFunc {
	return func(c echo.Context) error {

		userInfo := c.Get(middleware.UserInfo).(*model.UserInfo)

		return c.JSON(http.StatusOK, userInfo)
	}

}
