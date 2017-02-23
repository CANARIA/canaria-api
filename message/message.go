package message

const (
	REGISTERD_MAILADDRESS  = "既にメールアドレスが登録されています"
	PREREGISTER_SUBJECT    = "【仮登録】会員登録用URLのお知らせ"
	PRE_REGISTER_MAIL_BODY = `Canariaにご登録ありがとうございます。
  24時間以内に下記のURLからご登録下さい。
  `
	DUPULICATE_ACCOUNT = "そのユーザー名は既に使われています"
	INVALIED_TOKEN     = "このURLは不正です。有効期限が切れた、または既に登録されたユーザーが存在する可能性があります"
	REGISTER_SUBJECT   = "【Canaria】会員登録完了のお知らせ"
	REGISTER_MAIL_BODY = `本登録が完了しました。
  -----------------------
  登録内容は以下の通りです。

  ユーザーID: ${user_name}
  登録メールアドレス: ${mailaddress}
  -----------------------

  このメールは送信専用メールアドレスです。
  ご返信頂いてもお答えできませんのでご了承ください。

  このメールに心当たりのない方はお手数ですが、以下のメールアドレスからご連絡ください。
  ***************************
  ここには署名が入ります。
  ***************************
  `
	INVALIED_LOGIN_CLAIM = "ユーザー名またはパスワードが間違っています"
	BAD_REQUEST          = "不正なリクエストです"
	SYSTEM_ERROR         = "システムエラーが発生しました"
)
