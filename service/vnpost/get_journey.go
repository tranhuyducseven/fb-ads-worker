package vnpost

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type OrderStatusHistory struct {
	TraceDate    string  `json:"traceDate"`
	StatusText   string  `json:"statusText"`
	StatusDetail string  `json:"statusDetail"`
	PostmanName  string  `json:"postmanName"`
	PostmanTel   string  `json:"postmanTel"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
}

type Journey struct {
	OrderStatusHistoryDtoList []OrderStatusHistory `json:"orderStatusHistoryDtoList"`
}

func GetJourney(code string) (*Journey, error) {

	url := VNPOST_HOST + "/OrderTemplate/historynew?itemCode=" + code

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", "eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJKRmNhcC1XZWJBcGkiLCJleHAiOjE3MDAyNzk0NTMsIm5iZiI6MTY5OTY3NDM1MywiaWF0IjoxNjk5Njc0MzUzLCJhaWQiOiJNWVZOUCIsInVpZCI6IjAxQzY1MkMyLTkwNjYtNEIwQy04MjAzLTlEMTg1QjM4MzBDNyIsInVmbiI6IlbFqSBOaMOgbiBCdXNpbmVzcyIsIm9yZyI6IkMwMDU3MTIyOTQiLCJkaWQiOiI1YmRjN2I1NGE1ZmViNTU0MTE3YWQyZmRkODUxMGM4YyIsImxjcCI6MTY1NjQ4NTg5NDAwMCwiZXhwaXJhdGlvbkRhdGUiOjkwLCJpc0VtcGxveWVlIjpmYWxzZSwib3duZXIiOiIwMUM2NTJDMi05MDY2LTRCMEMtODIwMy05RDE4NUIzODMwQzciLCJwaG9uZU51bWJlciI6Iis4NDk2OTAzMjI0NiIsImFkZHIiOiJDxIJOIEE5IMSQxq_hu5xORyBT4buQIDIyIiwicHJvdiI6IjcwIiwiZGlzdCI6IjcxMzAiLCJjb21tIjoiNzEzMTUiLCJvcyI6IldFQiJ9.nQG_AMNitUi0SZLDgiW0cyWrzzmI2VkyVWI1y7gcNDLH55DWKLg8rmKjb7c9gIgPHu8viXeA5UYkosFh6n0KjA")
	req.Header.Add("Capikey", "19001111")

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

	var data *Journey
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
