package tokentable

type Token struct {
	ID string

	AuthCode string

	AccessToken           string
	AccessTokenExpireTime int64

	RefreshToken           string
	RefreshTokenExpireTime int64

	AppId          string
	AppSecret      string
	AppAccessToken string
	AppExpireTime  int64
}
