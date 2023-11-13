package vnpost

import (
	"bytes"
	"encoding/json"
	"errors"
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

	req.Header.Add("authorization", "eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJKRmNhcC1XZWJBcGkiLCJleHAiOjE3MDAyMTMzNTQsIm5iZiI6MTY5OTYwODI1NCwiaWF0IjoxNjk5NjA4MjU0LCJhaWQiOiJNWVZOUCIsInVpZCI6IjAxQzY1MkMyLTkwNjYtNEIwQy04MjAzLTlEMTg1QjM4MzBDNyIsInVmbiI6IlbFqSBOaMOgbiBCdXNpbmVzcyIsIm9yZyI6IkMwMDAyMTg5NzAiLCJkaWQiOiI1YmRjN2I1NGE1ZmViNTU0MTE3YWQyZmRkODUxMGM4YyIsImxjcCI6MTY1NjQ4NTg5NDAwMCwiZXhwaXJhdGlvbkRhdGUiOjkwLCJpc0VtcGxveWVlIjpmYWxzZSwib3duZXIiOiIwMUM2NTJDMi05MDY2LTRCMEMtODIwMy05RDE4NUIzODMwQzciLCJwaG9uZU51bWJlciI6Iis4NDk2OTAzMjI0NiIsImFkZHIiOiJDxIJOIEE5IMSQxq_hu5xORyBT4buQIDIyIiwicHJvdiI6IjcwIiwiZGlzdCI6IjcxMzAiLCJjb21tIjoiNzEzMTUiLCJvcyI6IldFQiJ9.LvSGwJcK1QfppOkrvv64cizz670M9J9ezKf63MPwx_4mh4OQ00L-HzUL7AGxBPil7p1_36Fl0Ypfj1f6xEG3TA")
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

	return nil, errors.New("cannot find order")

}
