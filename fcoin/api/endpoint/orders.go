package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func CreateOrderLimit(c *Configure, opts ...ParamsOption) (ret string, err error) {
	order := setParams(opts)
	order.Type = "limit"
	baseURL := c.getUrl("orders", "CreateOrderLimit")
	values := url.Values{}
	// alphabet order
	values.Add("type", order.Type)
	values.Add("side", order.Side)
	values.Add("symbol", order.Symbol)
	values.Add("price", order.Price)
	values.Add("amount", order.Amount)
	reader := bytes.NewBuffer(validJSONBody(values))
	// alphabet order
	validPayload := map[string]string{"amount": order.Amount, "side": order.Side, "symbol": order.Symbol, "price": order.Price, "type": order.Type}
	ret, err = apiConfig(c).Post(baseURL, reader, validPayload, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func OrderList(c *Configure, opts ...ParamsOption) (ret string, err error) {
	order := setParams(opts)
	baseURL := c.getUrl("orders", "OrderList")
	values := url.Values{}
	values.Add("symbol", order.Symbol)
	values.Add("states", order.States)
	values.Add("limit", adjustedPerPage(order))
	values.Add("before", order.PageBefore)
	values.Add("after", order.PageAfter)
	query := values.Encode()
	// alphabet order
	validPayload := map[string]string{"after": order.PageAfter, "before": order.PageBefore, "limit": adjustedPerPage(order), "states": order.States, "symbol": order.Symbol}
	ret, err = apiConfig(c).Get(baseURL, query, validPayload, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ReferenceOrder(c *Configure, opts ...ParamsOption) (ret string, err error) {
	order := setParams(opts)
	url := c.getUrl("orders", "ReferenceOrder") + "/" + order.OrderId
	ret, err = apiConfig(c).Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func CancelOrder(c *Configure, opts ...ParamsOption) (ret string, err error) {
	order := setParams(opts)
	url := c.getUrl("orders", "CancelOrder") + "/" + order.OrderId + "/submit-cancel"
	ret, err = apiConfig(c).Post(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func OrderMatchResults(c *Configure, opts ...ParamsOption) (ret string, err error) {
	order := setParams(opts)
	url := c.getUrl("orders", "OrderMatchResults") + "/" + order.OrderId + "/match-results"
	ret, err = apiConfig(c).Get(url, nil, nil, true)
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

func adjustedPerPage(pa *Params) (ret string) {
	if pa.PerPage != "" {
		ret = pa.PerPage
	} else {
		ret = "20"
	}
	return
}
