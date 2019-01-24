package src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func checkStatusCode(res *http.Response, statusCode int, action string) error {
	if res.StatusCode != statusCode {
		body, _ := ioutil.ReadAll(res.Body)
		m := fmt.Sprintf("Unexptected status code %d when %s (response: %s)", res.StatusCode, action, string(body))
		glog.Error(m)
		return errors.New(m)
	}
	return nil
}
