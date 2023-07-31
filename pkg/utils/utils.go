package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func checkEmptyFields(data interface{}) error {
	value := reflect.ValueOf(data)

	// Pastikan data adalah sebuah struct
	if value.Kind() != reflect.Struct {
		return errors.New("data harus berupa struct")
	}

	// Loop melalui setiap field dari struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		// Periksa apakah field memiliki nilai kosong
		if field.Interface() == reflect.Zero(field.Type()).Interface() {
			return errors.New("There is empty field")
		}
	}

	return nil
}
