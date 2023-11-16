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

	req.Header.Add("authorization", "eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJKRmNhcC1XZWJBcGkiLCJleHAiOjE3MDAyNzk0NTMsIm5iZiI6MTY5OTY3NDM1MywiaWF0IjoxNjk5Njc0MzUzLCJhaWQiOiJNWVZOUCIsInVpZCI6IjAxQzY1MkMyLTkwNjYtNEIwQy04MjAzLTlEMTg1QjM4MzBDNyIsInVmbiI6IlbFqSBOaMOgbiBCdXNpbmVzcyIsIm9yZyI6IkMwMDU3MTIyOTQiLCJkaWQiOiI1YmRjN2I1NGE1ZmViNTU0MTE3YWQyZmRkODUxMGM4YyIsImxjcCI6MTY1NjQ4NTg5NDAwMCwiZXhwaXJhdGlvbkRhdGUiOjkwLCJpc0VtcGxveWVlIjpmYWxzZSwib3duZXIiOiIwMUM2NTJDMi05MDY2LTRCMEMtODIwMy05RDE4NUIzODMwQzciLCJwaG9uZU51bWJlciI6Iis4NDk2OTAzMjI0NiIsImFkZHIiOiJDxIJOIEE5IMSQxq_hu5xORyBT4buQIDIyIiwicHJvdiI6IjcwIiwiZGlzdCI6IjcxMzAiLCJjb21tIjoiNzEzMTUiLCJvcyI6IldFQiJ9.nQG_AMNitUi0SZLDgiW0cyWrzzmI2VkyVWI1y7gcNDLH55DWKLg8rmKjb7c9gIgPHu8viXeA5UYkosFh6n0KjA")
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
