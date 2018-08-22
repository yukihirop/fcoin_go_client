package api

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Configure) Get(url string, query interface{}, payload map[string]string, isAuthorize bool) (ret string, err error) {
	ret, err = c.Request("GET", url, query, payload, isAuthorize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (c *Configure) Post(url string, reader interface{}, payload map[string]string, isAuthorize bool) (ret string, err error) {
	ret, err = c.Request("POST", url, reader, payload, isAuthorize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (c *Configure) Request(httpMethod string, url string, query_or_reader interface{}, payload map[string]string, isAuthorize bool) (ret string, err error) {
	var req *http.Request
	switch httpMethod {
	case "GET":
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		// set query parametter
		query, _ := query_or_reader.(string)
		req.URL.RawQuery = query
	case "POST":
		body, _ := query_or_reader.(io.Reader)
		req, err = http.NewRequest("POST", url, body)
		if err != nil {
			fmt.Println(err)
		}
		// Content-Type (json)
		req.Header.Set("Content-Type", "application/json")
	}

	// Authorization
	if isAuthorize {
		c.authorize(req, httpMethod, url, payload)
	}

	ret, _ = Connect(req)
	return
}

func (c *Configure) authorize(req *http.Request, httpMethod string, url string, payload map[string]string) {
	auth := NewAuthorization(httpMethod, url, c.ApiKey, c.SecretKey, payload)
	originalHeaders := auth.OriginalHeaders()
	for key, value := range originalHeaders {
		req.Header.Add(key, value)
	}
}
