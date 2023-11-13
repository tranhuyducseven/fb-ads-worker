package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"
)

func ConvertMapToStruct(m map[string]interface{}, s interface{}) error {
	stValue := reflect.ValueOf(s).Elem()
	sType := stValue.Type()
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		if value, ok := m[field.Name]; ok {
			stValue.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return nil
}

func ConvertTemplate(file string, executeData interface{}, subject string) []byte {
	var body bytes.Buffer
	t, _ := template.ParseFiles(file)
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))

	t.Execute(&body, executeData)

	return body.Bytes()
}

type MailCreateUserTemplate struct {
	Password string
	Email    string
}

type MailDepositTemplate struct {
	Amount string
}
