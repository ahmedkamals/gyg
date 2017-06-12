package http

import (
	"io/ioutil"
	"net/http"
	"gyg/middleware/communication"
	"strings"
)

type Response struct {
	statusCode   int
	headers      map[string]string
	body         string
	baseResponse *communication.Response
}

func NewHttpResponseFromNative(resp *http.Response) (*Response, error) {
	headers := make(map[string]string, 5)

	for key, values := range resp.Header {
		// Converting the array of bytes to string
		headers[key] = strings.Join(values, ",")
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	bodyAsString := string(body)

	var isSuccessful bool

	if 300 > resp.StatusCode {
		isSuccessful = true
	}

	return &Response{
		statusCode: resp.StatusCode,
		headers:    headers,
		body:       bodyAsString,
		baseResponse: communication.NewResponse(
			bodyAsString,
			headers,
			isSuccessful,
		),
	}, nil
}

func (self *Response) StatusCode() int {
	return self.statusCode
}

func (self *Response) Headers() map[string]string {
	return self.headers
}

func (self *Response) Body() string {
	return self.body
}

func (self *Response) IsSuccessful() bool {
	return self.baseResponse.IsSuccessful()
}

func (self *Response) toBaseResponse() *communication.Response {
	return self.baseResponse
}
