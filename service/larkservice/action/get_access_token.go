package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetAccessTokenResponse struct {
	AuthData AuthData `json:"data"`
	Code     int      `json:"code"`
	Msg      string   `json:"msg"`
}

func GetAccessToken(code, token string) (*GetAccessTokenResponse, error) {

	url := "https://open.larksuite.com/open-apis/authen/v1/access_token"

	client := &http.Client{}

	requestData := map[string]string{
		"grant_type": "authorization_code",
		"code":       code,
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

	data := GetAccessTokenResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
