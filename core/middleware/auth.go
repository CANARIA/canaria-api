package middleware

import (
	"fmt"
	"net/http"

	"github.com/CANARIA/canaria-api/core/config"
	"github.com/CANARIA/canaria-api/core/message"
	"github.com/CANARIA/canaria-api/core/model"
	"github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func AuthFilter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			accessToken := c.Request().Header.Get("access_token")
			claimedJwt := c.Request().Header.Get("Authorization")

			if accessToken == "" && claimedJwt == "" {
				logrus.Info("failed: Unauthorized")
				return echo.NewHTTPError(http.StatusUnauthorized, message.BAD_REQUEST)
			}

			userInfo := model.UserInfo{}
			token, err := jwt.ParseWithClaims(claimedJwt, &userInfo, func(token *jwt.Token) (interface{}, error) {
				// Make sure token's signature wasn't changed
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected siging method")
				}
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
				return c.JSON(http.StatusUnauthorized, "認証リクエストが不正です")
			}

			// contextにユーザー情報を含む
			c.Set(config.UserInfo, &userInfo)

			return next(c)

		})
	}
}
