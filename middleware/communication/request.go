package communication

import (
	"strconv"
	"strings"
)

type (
	RequestAware interface {
		GetPayload() []byte
		GetMeta() map[string]string
		String() string
	}

	Request struct {
		payload []byte
		meta    map[string]interface{}
	}
)

func NewRequest(body []byte, meta map[string]interface{}) *Request {
	return &Request{body, meta}
}

func (req *Request) GetPayload() []byte {
	return req.payload
}

func (req *Request) SetMeta(meta map[string]interface{}) *Request {
	req.meta = meta
	return req
}

func (req *Request) GetMeta() map[string]string {
	copy := make(map[string]string)
	for key, item := range req.meta {

		switch value := item.(type) {
		case string:
			copy[key] = value
			break
		case int:
			copy[key] = strconv.Itoa(value)
			break
		case map[string]string:
			copy[key] = BuildParams(value)
			break
		}
	}

	return copy
}

func (req *Request) String() string {
	return BuildParams(req.GetMeta())
}

func BuildParams(inputParams map[string]string) string {
	params := []string{}
	for param, value := range inputParams {
		params = append(params, param+"="+value)
	}

	return strings.Join(params, "&")
}
