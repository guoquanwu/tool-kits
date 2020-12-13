package http

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// http: send httpPost request
// every request will create a http client, this is not a effective way, just for low frenquency call
func Post(method, url, username, passwd string, postBody string, timeout int64) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	req, err := http.NewRequest(method, url, strings.NewReader(postBody))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, err
}
