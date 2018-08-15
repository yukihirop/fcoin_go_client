package api

import (
	"fmt"
	"io"
	"net/http"
)

func (c *APIConfigure) Get(url string, body io.Reader, payload map[string]string, isAuthorize bool) (ret string, err error) {
	ret, err = c.Request("GET", url, body, payload, isAuthorize)
	return
}

func (c *APIConfigure) Post(url string, body io.Reader, payload map[string]string, isAuthorize bool) (ret string, err error) {
	ret, err = c.Request("POST", url, body, payload, isAuthorize)
	return
}

func (c *APIConfigure) Request(httpMethod string, url string, body io.Reader, payload map[string]string, isAuthorize bool) (ret string, err error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Content-Type (URL encode)
	req.Header.Set("Content-Type", "application/json")

	// Authorization
	if isAuthorize {
		c.authorize(req, httpMethod, url, payload)
	}

	ret, _ = Connect(req)
	return
}

func (c *APIConfigure) authorize(req *http.Request, httpMethod string, url string, payload map[string]string) {
	auth := NewAuthorization(httpMethod, url, c.ApiKey, c.SecretKey, payload)
	originalHeaders := auth.OriginalHeaders()
	for key, value := range originalHeaders {
		req.Header.Add(key, value)
	}
}
