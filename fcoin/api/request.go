package api

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) (ret string, err error) {
	ret, err = Request("GET", url, nil)
	return
}

func Post(url string, body *io.Reader) (ret string, err error) {
	ret, err = Request("POST", url, body)
	return
}

func Request(httpMethod string, url string, body *io.Reader) (ret string, err error) {
	req, err := http.NewRequest(httpMethod, url, nil)
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
