package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const contentType = "application/json; charset=utf-8"

var (
	ErrNoBody             = errors.New("request has no json body")
	ErrContentTypeNotJson = errors.New("request does not provide correct content type")
)

func WriteJson(w http.ResponseWriter, data any) error {
	w.Header().Add("Content-Type", contentType)

	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}

func WriteError(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, "err: %s", err) // Here we are indeed ignoring errors
}

func GetJson[T any](r *http.Request) (*T, error) {
	ct := r.Header.Get("Content-Type")

	if !strings.Contains(ct, "application/json") {
		return nil, ErrContentTypeNotJson
	}

	if r.ContentLength == 0 {
		return nil, ErrNoBody
	}

	decoder := json.NewDecoder(r.Body)

	val := new(T)
	err := decoder.Decode(val)

	return val, err
}
