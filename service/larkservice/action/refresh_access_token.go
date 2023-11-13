package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RefreshAccessTokenResponse struct {
	AuthData AuthData `json:"data"`
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
}
type AuthData struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	ExpiresIn        int64  `json:"expires_in"`
}

func RefreshAccessToken(refresh_token, token string) (*RefreshAccessTokenResponse, error) {

	url := "https://open.larksuite.com/open-apis/authen/v1/refresh_access_token"

	client := &http.Client{}

	requestData := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": refresh_token,
	}

	postBody, _ := json.Marshal(requestData)
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := RefreshAccessTokenResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
