package mock

import (
	"fcoin_go_client/fcoin"
	"fcoin_go_client/fcoin/api"
	"fcoin_go_client/fcoin/api/endpoint"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	govcr "gopkg.in/seborama/govcr.v2"
)

type Mock struct {
	endPoint   string
	cassetName string
	APIKey     string
	SecretKey  string
	Payload    map[string]string
}

type MockAPI interface {

	//public
	PublicServerTime() (string, error)
	PublicCurrencies() (string, error)
	PublicSymbols() (string, error)

	//market
	MarketTicker(...MockEndpointOption) (string, error)
	MarketDepth(...MockEndpointOption) (string, error)
	MarketTrades(...MockEndpointOption) (string, error)
	MarketCandles(...MockEndpointOption) (string, error)

	//orders
	CreateOrderLimit(...MockEndpointOption) (string, error)
}

func NewMockAPI(cassetName string) MockAPI {
	api := fcoin.NewAPI()
	return &Mock{
		endPoint:   api.GetEndPoint(),
		cassetName: cassetName,
		APIKey:     os.Getenv("FCOIN_API_KEY"),
		SecretKey:  os.Getenv("FCOIN_SECRET_KEY"),
	}
}

func (m *Mock) url(endPoint string, methodName string) (ret string) {
	path := endpoint.GetPath(endPoint, methodName)
	ret = m.endPoint + path
	return
}

func (m *Mock) Get(url string) (ret string, err error) {
	ret, err = m.Request("GET", url, nil, nil)
	return
}

func (m *Mock) Post(url string, body io.Reader, payload map[string]string) (ret string, err error) {
	ret, err = m.Request("POST", url, body, payload)
	return
}

func (m *Mock) Request(httpMethod string, url string, body io.Reader, payload map[string]string) (ret string, err error) {
	var req *http.Request
	switch httpMethod {
	case "GET":
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
	case "POST":
		// TODO: vcrを使うようにする
		req, err = http.NewRequest("POST", url, body)
		if err != nil {
			fmt.Println(err)
		}
		// Content-Type (json)
		req.Header.Set("Content-Type", "application/json")
	}

	// authorize
	m.authorize(req, httpMethod, url, payload)
	// vcr
	vcr := govcr.NewVCR(m.cassetName,
		&govcr.VCRConfig{
			ExcludeHeaderFunc: func(key string) bool {
				return strings.EqualFold(key, "Fc-Access-Key") && strings.EqualFold(key, "Fc-Access-Signature") && strings.EqualFold(key, "Fc-Access-Timestamp")
			},
			// do not working
			RequestFilterFunc: func(header http.Header, body []byte) (*http.Header, *[]byte) {
				header.Set("Fc-Access-Key", "<Fc-Access-Key>")
				header.Set("Fc-Access-Signature", "<Fc-Access-Signature>")
				header.Set("Fc-Access-Timestamp", "<Fc-Access-Timestamp>")
				return &header, &body
			},
			Logging: true,
		})
	resp, err := vcr.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	ret, err = Execute(resp)
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

func (m *Mock) authorize(req *http.Request, httpMethod string, url string, payload map[string]string) {
	auth := api.NewAuthorization(httpMethod, url, m.APIKey, m.SecretKey, payload)
	originalHeaders := auth.OriginalHeaders()
	for key, value := range originalHeaders {
		req.Header.Add(key, value)
	}
}
