package api

import (
	"fmt"

	"net/http"

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
		// たとえばリクエストのヘダーからトークン文字列を受け取ったとする
		tokenstring := CreateTokenString()
		accessToken := c.Request().Header.Get("access_token")
		claimedJwt := c.Request().Header.Get("jwt")

		fmt.Println(tokenstring)
		// サーバだけが知り得るSecretでこれをParseする

		// 別例, jwt.StandardClaimsを満たすstructに直接decodeさせることもできる
		userInfo := model.UserInfo{}
		token, err := jwt.ParseWithClaims(claimedJwt, &userInfo, func(token *jwt.Token) (interface{}, error) {
			return []byte("ServerSecretKey"), nil
		})
		fmt.Println(token.Valid, userInfo, err)
		fmt.Println("Token valid ? : ", token.Valid)
		fmt.Println("User's Info : ", userInfo)

		// トークンのパースに失敗
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "セッションの有効期限が切れた可能性があります。再度ログインし直してください。")
		}

		// リクエストのアクセストークンとJWTに含まれたアクセストークンが一致しなかったら認証失敗
		if accessToken != userInfo.AccessToken {
			return c.JSON(http.StatusUnauthorized, "認証リクエストが不正です。")
		}

		// ここまでがcontextに含まれる内容、ここからAPIの本処理

		return c.JSON(http.StatusOK, "人気のタグ一覧が入るよ")
	}

}
