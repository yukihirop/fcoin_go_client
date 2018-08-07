package mock

import (
	"fcoin_go_client/fcoin"
	"fcoin_go_client/fcoin/endpoint"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	govcr "gopkg.in/seborama/govcr.v2"
)

type Mock struct {
	endPoint   string
	cassetName string
}

func NewMockAPI(cassetName string) (m *Mock) {
	api := fcoin.NewAPI()
	m = &Mock{endPoint: api.GetEndPoint(), cassetName: cassetName}
	return
}

func (m *Mock) url(endPoint string, methodName string) (ret string) {
	path := endpoint.GetPath(endPoint, methodName)
	ret = m.endPoint + path
	return
}

func (m *Mock) Get(url string) (ret string, err error) {
	ret, err = m.Request("GET", url, nil)
	return
}

func (m *Mock) Post(url string, body io.Reader) (ret string, err error) {
	ret, err = m.Request("POST", url, body)
	return
}

func (m *Mock) Request(httpMethod string, url string, body io.Reader) (ret string, err error) {
	vcr := govcr.NewVCR(m.cassetName, nil)
	switch httpMethod {
	case "GET":
		resp, err := vcr.Client.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		ret, err = Execute(resp)
	case "POST":
		resp, err := vcr.Client.Post(url, "application/json", body)
		if err != nil {
			fmt.Println(err)
		}
		ret, err = Execute(resp)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
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
