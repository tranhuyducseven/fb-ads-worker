package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetBaseListRecordResquest struct {
	AppId     string
	TableId   string
	ViewId    string
	PageToken string
}

type GetBaseListRecordResponse struct {
	Records Records `json:"data"`
	Code    int     `json:"code"`
	Msg     string  `json:"msg"`
}

type Records struct {
	HasMore   bool     `json:"has_more"`
	Items     []Record `json:"items"`
	PageToken string   `json:"page_token"`
	Total     int      `json:"total"`
}

type Record struct {
	Id       string      `json:"id"`
	RecordId string      `json:"record_id"`
	Fields   interface{} `json:"fields"`
}

func GetBaseListRecord(request GetBaseListRecordResquest, access_token string) (*GetBaseListRecordResponse, error) {

	url := "https://open.larksuite.com/open-apis/bitable/v1/apps/%v/tables/%v/records"

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(url, request.AppId, request.TableId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", access_token))

	q := req.URL.Query()
	q.Add("view_id", request.ViewId)
	q.Add("page_token", request.PageToken)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := GetBaseListRecordResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
