package larkservice

import (
	larkaction "data-pipeline/service/larkservice/action"
	"errors"
)

var access_token string
var refresh_token string

func GetAccessToken(code, token string) error {

	res, err := larkaction.GetAccessToken(code, token)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New("error can't get access token")
	}

	access_token = res.AuthData.AccessToken
	refresh_token = res.AuthData.RefreshToken

	return nil
}

func RefreshAccessToken(token string) error {
	res, err := larkaction.RefreshAccessToken(refresh_token, token)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New("error can't get access token")
	}

	access_token = res.AuthData.AccessToken
	refresh_token = res.AuthData.RefreshToken
	
	return nil
}
