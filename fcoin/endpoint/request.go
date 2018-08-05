package endpoint

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) (int, error) {
	return Request("GET", url, nil)
}

func Post(url string, body *io.Reader) (int, error) {
	return Request("POST", url, body)
}

func Request(httpMethod string, url string, body *io.Reader) (int, error) {
	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return fmt.Println(err)
	}
	// req.URL.RawQuery = url.values.Encode()

	// Add Authorization Header
	// req.Header.Add

	return Connect(req)
}
