package api

import (
	"net/http"

	"github.com/CANARIA/canaria-api/mail"
	"github.com/CANARIA/canaria-api/message"
	"github.com/CANARIA/canaria-api/model"
	"github.com/CANARIA/canaria-api/util"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

func PreRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		var account model.Account

		preJson := new(model.PreRegister)
		if err := c.Bind(preJson); err != nil {
			return err
		}

		tx := c.Get("Tx").(*dbr.Tx)

		tx.Select("*").
			From("accounts").
			Where("mailaddress = ?", preJson.MailAddress).
			Load(&account)

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

		tx := c.Get("Tx").(*dbr.Tx)

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

		tx := c.Get("Tx").(*dbr.Tx)

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
