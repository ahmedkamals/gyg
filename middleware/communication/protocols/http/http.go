package http

import (
	"bytes"
	"net/http"
	"gyg/middleware/communication"
	"strings"
)

type (
	Protocol struct {
		client *http.Transport
	}
)

func NewProtocol(client *http.Transport) *Protocol {
	return &Protocol{client}
}

func (self *Protocol) Client() *http.Transport {
	return self.client
}

func (self *Protocol) Send(req communication.RequestAware) (*communication.Response, error) {
	httpRequest, err := NewHttpRequest(req)
	if err != nil {
		return nil, err
	}

	nativeRequest, err := http.NewRequest(
		httpRequest.GetMethod(),
		self.GetUrl(httpRequest),
		bytes.NewBuffer(httpRequest.GetBody()),
	)
	if err != nil {
		return nil, err
	}

	// Setting headers
	for key, value := range httpRequest.headers {
		nativeRequest.Header.Add(key, value)
	}

	nativeResponse, err := self.Client().RoundTrip(nativeRequest)
	if err != nil {
		return nil, err
	}

	httpResponse, err := NewHttpResponseFromNative(nativeResponse)
	if err != nil {
		return nil, err
	}

	// Making sure the buffer will be closed.
	nativeResponse.Body.Close()

	return httpResponse.toBaseResponse(), nil
}

func (self *Protocol) GetUrl(httpRequest *Request) string {
	return strings.Join([]string{httpRequest.GetProtocol(), httpRequest.GetHost(), httpRequest.GetPort(), httpRequest.GetPath()}, "") + "?" + httpRequest.GetParams()
}
