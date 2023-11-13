package nhanh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func SearchProduct(page, limit int) (*SearchProductResponse,error) {
	url := "https://open.nhanh.vn/api/product/search"

	requestData := map[string]string{
		"version":     os.Getenv("NHANH_VERSION"),
		"appId":       os.Getenv("NHANH_APP_ID"),
		"businessId":  os.Getenv("NHANH_BUSINESS_ID"),
		"accessToken": os.Getenv("NHANH_ACCESS_TOKEN"),
		"data":        fmt.Sprintf(`{"page": %v, "icpp": %v}`, page, limit),
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for key, value := range requestData {
		w.WriteField(key, value)
	}

	w.Close()

	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil,err
	}

	data := SearchProductResponse{}
	if err := json.Unmarshal(resBody, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
