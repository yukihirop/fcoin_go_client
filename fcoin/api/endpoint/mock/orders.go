package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func (m *Mock) CreateOrderLimit(opts ...MockEndpointOption) (ret string, err error) {
	order := setMockEndpoint(opts)
	order.Type = "limit"
	baseURL := "https://api.fcoin.com/v2/orders"
	values := url.Values{}
	values.Add("type", order.Type)
	values.Add("side", order.Side)
	values.Add("symbol", order.Symbol)
	values.Add("price", order.Price)
	values.Add("amount", order.Amount)

	reader := bytes.NewBuffer(validJSONBody(values))
	// alphabet order
	validPayload := map[string]string{"amount": order.Amount, "side": order.Side, "symbol": order.Symbol, "price": order.Price, "type": order.Type}
	ret, err = m.Post(baseURL, reader, validPayload, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) OrderList(opts ...MockEndpointOption) (ret string, err error) {
	order := setMockEndpoint(opts)
	baseURL := "https://api.fcoin.com/v2/orders"
	values := url.Values{}
	values.Add("symbol", order.Symbol)
	values.Add("states", order.States)
	values.Add("limit", adjustedPerPage(order))
	values.Add("before", order.PageBefore)
	values.Add("after", order.PageAfter)

	query := values.Encode()
	// alphabet order
	validPayload := map[string]string{"after": order.PageAfter, "before": order.PageBefore, "limit": adjustedPerPage(order), "states": order.States, "symbol": order.Symbol}
	ret, err = m.Get(baseURL, query, validPayload, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) ReferenceOrder(opts ...MockEndpointOption) (ret string, err error) {
	order := setMockEndpoint(opts)
	url := "https://api.fcoin.com/v2/orders/" + order.OrderId
	ret, err = m.Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) CancelOrder(opts ...MockEndpointOption) (ret string, err error) {
	order := setMockEndpoint(opts)
	url := "https://api.fcoin.com/v2/orders/" + order.OrderId + "/submit-cancel"
	ret, err = m.Post(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) OrderMatchResults(opts ...MockEndpointOption) (ret string, err error) {
	order := setMockEndpoint(opts)
	url := "https://api.fcoin.com/v2/orders/" + order.OrderId + "/match-results"
	ret, err = m.Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func validJSONBody(values map[string][]string) (ret []byte) {
	var params = make(map[string]string, 1)
	for k, v := range values {
		params[k] = v[0]
	}
	ret, _ = json.Marshal(params)
	return
}

func adjustedPerPage(order *MockEndpoint) (ret string) {
	if order.PerPage != "" {
		ret = order.PerPage
	} else {
		ret = "20"
	}
	return
}
