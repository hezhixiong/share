package ctr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetQuery(r *http.Request, key string, defaultStr string) string {
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys) < 1 {
		return defaultStr
	}
	return keys[0]
}

func GetQueryPage(r *http.Request) int {
	page := GetQuery(r, "page", strconv.Itoa(PAGE_DEFAULT))
	result, err := strconv.Atoi(page)
	if err != nil || result < 1 {
		result = PAGE_DEFAULT
	}
	return result
}

func GetQuerySize(r *http.Request) int {
	size := GetQuery(r, "size", strconv.Itoa(PAGE_SIZE_DEFAULT))
	result, err := strconv.Atoi(size)
	if err != nil || result < PAGE_SIZE_MIN || result > PAGE_SIZE_MAX {
		result = PAGE_SIZE_DEFAULT
	}
	return result
}

func BindJson(r *http.Request, object interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := r.Body.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(body, &object); err != nil {
		return err
	}

	return nil
}
