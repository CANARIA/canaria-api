package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CANARIA/canaria-api/core/mail"
	"github.com/CANARIA/canaria-api/core/message"
	"github.com/CANARIA/canaria-api/core/model"
	"github.com/CANARIA/canaria-api/core/util"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/CANARIA/canaria-api/core/config"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

// 仮登録
func PreRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		var account model.Account
		ctx := appengine.NewContext(c.Request())

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
		mail.Send(ctx)

		return c.JSON(http.StatusOK, "Ok")
	}
}

// トークンチェック
func CheckToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		checkTokenJson := new(model.CheckToken)
		if err := c.Bind(checkTokenJson); err != nil {
			return err
		}
		auth := model.Auth{UrlToken: &checkTokenJson.UrlToken}

		tx := c.Get("Tx").(*gorm.DB)

		preAccountDao := model.PreAccountDaoFactory(tx)
		res, err := preAccountDao.ValidPreAccountToken(&auth)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.INVALIED_TOKEN)
		}

		if res == nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.DUPULICATE_ACCOUNT)
		}

		return c.JSON(http.StatusOK, "ok")
	}
}

/*
  ユーザー登録
*/
func AuthRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := appengine.NewContext(c.Request())

		authJson := new(model.AuthRegister)
		if err := c.Bind(authJson); err != nil {
			return err
		}
		auth := model.Auth{UrlToken: &authJson.UrlToken}

		tx := c.Get("Tx").(*gorm.DB)

		preAccountDao := model.PreAccountDaoFactory(tx)
		// トークンチェック
		res, err := preAccountDao.ValidPreAccountToken(&auth)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.INVALIED_TOKEN)
		}

		// TODO: トランザクション

		// アカウントの作成
		accountDao := model.AccountDaoFactory(tx)
		account := model.AccountImpl(authJson, res)

		if err := accountDao.Create(account); err != nil {
			fmt.Errorf(err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, message.DUPULICATE_ACCOUNT)
		}

		// プロフィールの作成
		profileDao := model.ProfileDaoFactory(tx)

		profile := profileDao.ProfileImpl(authJson, account)

		if err := profileDao.Create(profile); err != nil {
			fmt.Errorf(err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, message.SYSTEM_ERROR)
		}

		// 仮登録情報の更新
		preAccount := model.BuildPreAccountEntity(&auth, res)
		if err := preAccountDao.AcctivateAccount(preAccount); err != nil {
			fmt.Errorf(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, message.SYSTEM_ERROR)
		}

		// TODO: Session発行
		// 登録確認メール送信
		mail := mail.BuildRegisteredMail(preAccount)
		mail.Send(ctx)

		return c.JSON(http.StatusOK, authJson)
	}
}

/*
  ログイン
*/
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userInfo model.UserInfo
		// headerから取得
		accessToken := c.Request().Header.Get("access_token")

		if accessToken == "" {
			return echo.NewHTTPError(http.StatusBadRequest, message.BAD_REQUEST)
		}

		loginClaim := new(model.LoginClaim)
		if err := c.Bind(loginClaim); err != nil {
			return err
		}
		tx := c.Get("Tx").(*gorm.DB)

		// userInfoにユーザー情報を詰めるため結合している
		res := tx.Select("*").
			Table("accounts a").
			Joins("LEFT JOIN profiles p ON a.user_id = p.user_id").
			Where("a.user_name = ?", loginClaim.UserName).
			Find(&userInfo)

		userInfo.AccessToken = accessToken
		fmt.Println("userInfo => ", userInfo)

		if res.Error != nil {
			fmt.Printf("LoginClaim{err=%s}", res.Error.Error())
			return echo.NewHTTPError(http.StatusBadRequest, message.INVALIED_LOGIN_CLAIM)
		}

		_, err := util.IsValidPassword(userInfo.Password, loginClaim.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.INVALIED_LOGIN_CLAIM)
		}

		// TODO: token生成は共通化する
		// JWTの有効期限設定
		standardClaim := jwt.StandardClaims{
			// 有効期限は720時間（30日後）
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
			Issuer:    "canaria.io",
		}
		userInfo.StandardClaims = standardClaim
		// アクセストークンとユーザ情報を一緒にトークン化
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userInfo)
		// Secretで文字列にする. このSecretはサーバだけが知っている
		tokenstring, err := token.SignedString([]byte("ServerSecretKey"))
		if err != nil {
			fmt.Println(err)
		}

		respUserInfo := ConvertToRespUserInfo(userInfo)

		c.Response().Header().Set("access_token", accessToken)
		c.Response().Header().Set("Authorization", tokenstring)

		return c.JSON(http.StatusOK, &respUserInfo)
	}
}

/*
認証チェック
*/
func CheckAuth() echo.HandlerFunc {
	return func(c echo.Context) error {

		/*
		 middlewareの認証フィルターが通れば以降の処理が走る
		*/

		jwt := c.Request().Header.Get("Authorization")

		userInfo := c.Get(config.UserInfo).(*model.UserInfo)
		respUserInfo := ConvertToRespUserInfo(*userInfo)

		c.Response().Header().Set("access_token", userInfo.AccessToken)
		c.Response().Header().Set("Authorization", jwt)

		return c.JSON(http.StatusOK, respUserInfo)
	}

}

// トークン付きユーザー情報からクライアントに必要な情報だけをコンバート
func ConvertToRespUserInfo(userInfoWithAccessToken model.UserInfo) *model.RespUserInfo {
	return &model.RespUserInfo{
		UserID:      userInfoWithAccessToken.UserID,
		UserName:    userInfoWithAccessToken.UserName,
		DisplayName: userInfoWithAccessToken.DisplayName,
		MailAddress: userInfoWithAccessToken.MailAddress,
		Avatar:      userInfoWithAccessToken.Avatar,
		Roll:        userInfoWithAccessToken.Roll,
	}
}
