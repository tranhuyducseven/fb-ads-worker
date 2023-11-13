package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UpdateBaseRecordResquest struct {
	AppId    string
	TableId  string
	RecordId string
	Fields   map[string]interface{}
}

type UpdateBaseRecordResponse struct {
}

func UpdateBaseRecord(request UpdateBaseRecordResquest, access_token string) error {

	url := "https://open.larksuite.com/open-apis/bitable/v1/apps/%v/tables/%v/records/%v"

	client := &http.Client{}

	requestData := map[string]interface{}{
		"fields": request.Fields,
	}

	postBody, _ := json.Marshal(requestData)
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("PUT", fmt.Sprintf(url, request.AppId, request.TableId, request.RecordId), responseBody)
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
