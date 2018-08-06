// https://stackoverflow.com/questions/14416275/error-cant-load-package-package-my-prog-found-packages-my-prog-and-main?rq=1
package endpoint

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Connect(req *http.Request) (ret string, err error) {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// https://qiita.com/jpshadowapps/items/463b2623209479adcd88
	defer resp.Body.Close()
	ret, _ = Execute(resp)
	return
}

func Execute(resp *http.Response) (ret string, err error) {
	b, err := ioutil.ReadAll(resp.Body)
	ret = string(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
