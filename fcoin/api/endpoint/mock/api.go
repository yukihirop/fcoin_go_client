package mock

import (
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
	OrderList(...MockEndpointOption) (string, error)
	ReferenceOrder(...MockEndpointOption) (string, error)
	CancelOrder(...MockEndpointOption) (string, error)
	OrderMatchResults(...MockEndpointOption) (string, error)

	//accounts
	AccountsBalance() (string, error)
}

func NewMockAPI(cassetName string) MockAPI {
	return &Mock{
		endPoint:   "https://api.fcoin.com/v2/",
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

func (m *Mock) Get(url string, query interface{}, payload map[string]string, isAuthorize bool) (ret string, err error) {
	ret, err = m.Request("GET", url, query, payload, isAuthorize)
	return
}

func (m *Mock) Post(url string, reader interface{}, payload map[string]string, isAuthorize bool) (ret string, err error) {
	ret, err = m.Request("POST", url, reader, payload, isAuthorize)
	return
}

func (m *Mock) Request(httpMethod string, url string, query_or_reader interface{}, payload map[string]string, isAuthorize bool) (ret string, err error) {
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

	// authorize
	if isAuthorize {
		m.authorize(req, httpMethod, url, payload)
	}

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
		return
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
	if err != nil {
		fmt.Println(err)
		return
	}
	ret = string(b)
	return
}

func (m *Mock) authorize(req *http.Request, httpMethod string, url string, payload map[string]string) {
	auth := api.NewAuthorization(httpMethod, url, m.APIKey, m.SecretKey, payload)
	originalHeaders := auth.OriginalHeaders()
	for key, value := range originalHeaders {
		req.Header.Add(key, value)
	}
}
