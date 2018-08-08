package mock

import (
	"fcoin_go_client/fcoin"
	"fcoin_go_client/fcoin/api"
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

type MockAPI interface {

	//public
	PublicServerTime() (string, error)
	PublicCurrencies() (string, error)
	PublicSymbols() (string, error)

	//market
	MarketTicker(...Option) (string, error)
	MarketDepth(...Option) (string, error)
	MarketTrades(...Option) (string, error)
	MarketCandles(...Option) (string, error)
}

func NewMockAPI(cassetName string) MockAPI {
	api := fcoin.NewAPI()
	return &Mock{endPoint: api.GetEndPoint(), cassetName: cassetName}
}

func (m *Mock) url(endPoint string, methodName string) (ret string) {
	path := api.GetPath(endPoint, methodName)
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
