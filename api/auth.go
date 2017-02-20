package api

import (
	"fmt"
	"net/http"

	"github.com/CANARIA/canaria-api/mail"
	"github.com/CANARIA/canaria-api/message"
	"github.com/CANARIA/canaria-api/model"
	"github.com/CANARIA/canaria-api/util"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func PreRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		var account model.Account

		preJson := new(model.PreRegister)
		if err := c.Bind(preJson); err != nil {
			return err
		}

		tx := c.Get("Tx").(*gorm.DB)

		tx.Select("*").
			Table("accounts").
			Where("mailaddress = ?", preJson.MailAddress).
			Find(&account)

		if (model.Account{}) != account {
			return echo.NewHTTPError(http.StatusBadRequest, message.REGISTERD_MAILADDRESS)
		}

		// トークンをセット
		token := util.GenerateToken()
		preAccount := model.PreAccountImpl(preJson, token)
		if err := preAccount.PreAccountCreate(tx); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		url := model.BuildRegisterUrl(token)

		// mailを送る
		mail := mail.BuildPreRegisterMail(*preAccount, url)
		mail.Send()

		return c.JSON(http.StatusOK, "Ok")
	}
}

func CheckToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		checkTokenJson := new(model.CheckToken)
		if err := c.Bind(checkTokenJson); err != nil {
			return err
		}
		auth := model.Auth{UrlToken: checkTokenJson.UrlToken}

		tx := c.Get("Tx").(*gorm.DB)

		if _, err := auth.ValidPreAccountToken(tx); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.INVALIED_TOKEN)
		}

		return c.JSON(http.StatusOK, "ok")
	}
}

func AuthRegister() echo.HandlerFunc {
	return func(c echo.Context) error {

		authJson := new(model.AuthRegister)
		if err := c.Bind(authJson); err != nil {
			return err
		}
		auth := model.Auth{UrlToken: authJson.UrlToken, MailAddress: authJson.MailAddress}

		tx := c.Get("Tx").(*gorm.DB)

		if _, err := auth.ValidPreAccountToken(tx); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.INVALIED_TOKEN)
		}

		// アカウントの作成
		account := model.AccountImpl(authJson)
		if err := account.AccountCreate(tx); err != nil {
			println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// 仮登録情報の更新
		if err := model.AcctivateAccount(tx, &auth); err != nil {
			println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// TODO: Session発行
		// 登録確認メール送信
		mail := mail.BuildRegisteredMail(auth)
		mail.Send()

		return c.JSON(http.StatusOK, authJson)
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userInfo model.UserInfo

		loginClaim := new(model.LoginClaim)
		if err := c.Bind(loginClaim); err != nil {
			return err
		}
		tx := c.Get("Tx").(*gorm.DB)

		res := tx.Select("a.user_id, a.user_name, p.display_name, a.mailaddress, p.avatar, a.roll").
			Table("accounts a").
			Joins("INNER JOIN profiles p ON a.user_id = p.user_id").
			Where("a.user_name = ? AND a.password = ?", loginClaim.UserName, loginClaim.Password).
			Find(&userInfo)

		fmt.Println("userInfo => ", userInfo)

		if res.Error != nil {
			fmt.Errorf("LoginClaim{err=%s}", res.Error.Error())
			return echo.NewHTTPError(http.StatusBadRequest, message.INVALIED_LOGIN_CLAIM)
		}

		// TODO: token生成は共通化する
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &userInfo)
		// Secretで文字列にする. このSecretはサーバだけが知っている
		tokenstring, err := token.SignedString([]byte("ServerSecretKey"))
		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(http.StatusOK, tokenstring)
	}
}
