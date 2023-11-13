package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AddBaseRecordResquest struct {
	AppId   string
	TableId string
	Fields  map[string]interface{}
}

type AddBaseRecordResponse struct {
}

func AddBaseRecord(request AddBaseRecordResquest, access_token string) error {

	url := "https://open.larksuite.com/open-apis/bitable/v1/apps/%v/tables/%v/records"

	client := &http.Client{}

	requestData := map[string]interface{}{
		"fields": request.Fields,
	}

	postBody, _ := json.Marshal(requestData)
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", fmt.Sprintf(url, request.AppId, request.TableId), responseBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", access_token))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	data := GetBaseListRecordResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	return nil
}
