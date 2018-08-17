package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fcoin_go_client/fcoin/api/utility"
	"strconv"
	"time"
)

type Authorization interface {
	OriginalHeaders() map[string]string
}

type Auth struct {
	HTTPMethod string
	fullURL    string
	Payload    map[string]string
	APIKey     string
	SecretKey  string
	Timestamp  string
}

func NewAuthorization(httpMethod, url, apiKey, secretKey string, payload map[string]string) Authorization {
	return &Auth{
		HTTPMethod: httpMethod,
		fullURL:    url,
		Payload:    payload,
		APIKey:     apiKey,
		SecretKey:  secretKey,
	}
}

func (a *Auth) OriginalHeaders() map[string]string {
	return map[string]string{
		"FC-ACCESS-KEY":       a.APIKey,
		"FC-ACCESS-SIGNATURE": a.encodedSignature(),
		"FC-ACCESS-TIMESTAMP": a.timestamp(),
	}
}

// http://y0m0r.hateblo.jp/entry/20140410/1397146035
func (a *Auth) encodedSignature() string {
	encSignature := base64.StdEncoding.EncodeToString([]byte(a.signature()))
	key := a.SecretKey
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(encSignature))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func (a *Auth) signature() (ret string) {
	httpMethod := a.HTTPMethod
	queryString := a.queryString()
	fullURL := a.fullURL
	timestamp := a.timestamp()

	if httpMethod == "GET" && queryString != "" {
		ret = httpMethod + fullURL + "?" + queryString + timestamp
	} else {
		ret = httpMethod + fullURL + timestamp + queryString
	}
	return
}

func (a *Auth) queryString() (ret string) {
	qs := utility.NewSortedQuery(a.Payload).String()
	if qs != "" {
		ret = qs
	} else {
		ret = ""
	}
	return
}

func (a *Auth) timestamp() (ret string) {
	ts := a.Timestamp
	if ts != "" {
		ret = ts
	} else {
		now := time.Now().UnixNano()
		// digits 13
		// https://api.fcoin.com/v2/public/server-time
		// {"status":0,"data":1534059195554}
		millis := now / 1000000
		ret = strconv.FormatInt(millis, 10)
	}
	return
}
