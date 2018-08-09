package api

import (
	"fmt"
	"io"
	"net/http"
)

func (c *APIConfigure) Get(url string) (ret string, err error) {
	ret, err = c.Request("GET", url, nil)
	return
}

func (c *APIConfigure) Post(url string, body io.Reader) (ret string, err error) {
	ret, err = c.Request("POST", url, body)
	return
}

func (c *APIConfigure) Request(httpMethod string, url string, body io.Reader) (ret string, err error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// req.URL.RawQuery = url.values.Encode()

	// Add Authorization Header
	// req.Header.Add

	ret, _ = Connect(req)
	return
}
