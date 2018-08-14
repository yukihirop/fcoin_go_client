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
	baseURL := m.url("orders", "CreateOrderLimit")
	values := url.Values{}
	// alphabet order
	values.Add("type", order.Type)
	values.Add("side", order.Side)
	values.Add("symbol", order.Symbol)
	values.Add("price", order.Price)
	values.Add("amount", order.Amount)

	reader := bytes.NewBuffer(validJSONBody(values))
	ret, err = m.Post(baseURL, reader, validPayload(order))
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func validPayload(o *MockEndpoint) map[string]string {
	return map[string]string{
		"amount": o.Amount,
		"side":   o.Side,
		"symbol": o.Symbol,
		"price":  o.Price,
		"type":   o.Type,
	}
}

func validJSONBody(values map[string][]string) (ret []byte) {
	var params = make(map[string]string, 1)
	for k, v := range values {
		params[k] = v[0]
	}
	ret, _ = json.Marshal(params)
	return
}
