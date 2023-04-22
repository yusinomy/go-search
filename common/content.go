package common

import (
	"io/ioutil"
	"net/http"
)

func conetent(r *http.Response) string {
	content, _ := ioutil.ReadAll(r.Body)
	return string(content)
}
