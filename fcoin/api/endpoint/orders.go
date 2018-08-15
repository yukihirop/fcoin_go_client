package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func CreateOrderLimit(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	order := setEndpoint(opts)
	order.Type = "limit"
	baseURL := GetPath("orders", "CreateOrderLimit")
	values := url.Values{}
	// alphabet order
	values.Add("type", order.Type)
	values.Add("side", order.Side)
	values.Add("symbol", order.Symbol)
	values.Add("price", order.Price)
	values.Add("amount", order.Amount)
	reader := bytes.NewBuffer(validJSONBody(values))
	ret, err = c.Post(baseURL, reader, validPayload(order), true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func OrderList(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	order := setEndpoint(opts)
	baseURL := GetPath("orders", "OrderList")
	values := url.Values{}
	values.Add("symbol", order.Symbol)
	values.Add("states", order.States)
	values.Add("limit", adjustedPerPage(order))
	values.Add("before", order.PageBefore)
	values.Add("after", order.PageAfter)

	query := values.Encode()
	// alphabet order
	validPayload := map[string]string{"after": order.PageAfter, "before": order.PageBefore, "limit": adjustedPerPage(order), "states": order.States, "symbol": order.Symbol}
	ret, err = c.Get(baseURL, query, validPayload, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ReferenceOrder(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	order := setEndpoint(opts)
	url := GetPath("orders", "ReferenceOrder") + "/" + order.OrderId
	ret, err = c.Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func CancelOrder(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	order := setEndpoint(opts)
	url := GetPath("orders", "CancelOrder") + "/" + order.OrderId + "/submit-cancel"
	ret, err = c.Post(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func OrderMatchResults(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	order := setEndpoint(opts)
	url := GetPath("orders", "OrderMatchResults") + "/" + order.OrderId + "/match-results"
	ret, err = c.Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func validPayload(o *Endpoint) map[string]string {
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

func adjustedPerPage(order *Endpoint) (ret string) {
	if order.PerPage != "" {
		ret = order.PerPage
	} else {
		ret = "20"
	}
	return
}
