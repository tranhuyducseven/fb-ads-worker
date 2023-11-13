package action

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetAppAccessTokenResponse struct {
	AppAccessToken string `json:"app_access_token"`
	Expire         int64  `json:"expire"`
	Code           int    `json:"code"`
}

func GetAppAccessToken(app_id, app_secret string) (*GetAppAccessTokenResponse, error) {
	url := "https://open.larksuite.com/open-apis/auth/v3/app_access_token/internal"

	client := &http.Client{}

	requestData := map[string]string{
		"app_id":     app_id,
		"app_secret": app_secret,
	}

	postBody, _ := json.Marshal(requestData)
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := GetAppAccessTokenResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil

}
