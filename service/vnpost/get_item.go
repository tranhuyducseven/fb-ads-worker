package vnpost

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ItemVNPost struct {
	ItemCode string `json:"itemCode"`
}

func GetItem(code string) (*ItemVNPost, error) {

	url := VNPOST_HOST + "/OrderHdr/searchAllByParam"

	requestData := map[string]interface{}{
		"orderType": "1",
		"orgCode": []string{
			"C005712294",
		},
		"searchValue": code,
	}

	jsonValue, _ := json.Marshal(requestData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", VNPOST_TOKEN)
	req.Header.Add("Capikey", "19001111")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data []*ItemVNPost
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		return data[0], nil
	}

	return nil, nil

}
