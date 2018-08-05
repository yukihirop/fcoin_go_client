// https://stackoverflow.com/questions/14416275/error-cant-load-package-package-my-prog-found-packages-my-prog-and-main?rq=1
package endpoint

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Connect(req *http.Request) (int, error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Println(err)
	}

	defer resp.Body.Close()
	return Execute(resp)
}

func Execute(resp *http.Response) (ret int, err error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Println(err)
	}
	return fmt.Println(string(b))
}
