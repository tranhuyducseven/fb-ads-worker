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

	req.Header.Add("authorization", VNPOST_TOKEN)
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
