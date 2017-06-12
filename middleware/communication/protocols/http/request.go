package http

import (
	"errors"
	"gyg/middleware/communication"
	"strings"
)

type (
	Request struct {
		headers                                    map[string]string
		method, protocol, host, port, path, params string
		body                                       []byte
		baseRequest                                communication.RequestAware
	}
)

const (
	METHOD   = "method"
	PROTOCOL = "protocol"
	HOST     = "host"
	PORT     = "port"
	PATH     = "path"
	PARAMS   = "params"
	GET      = "GET"
	POST     = "POST"
	PUT      = "PUT"
	DELETE   = "DELETE"
)

func NewHttpRequest(req communication.RequestAware) (*Request, error) {
	var method, protocol, host, port, path, params string
	meta := req.GetMeta()

	for key, value := range meta {
		switch key {
		case METHOD:
			method = value
			delete(meta, key)
			break
		case PROTOCOL:
			protocol = value
			delete(meta, key)
			break
		case HOST:
			host = value
			delete(meta, key)
			break
		case PORT:
			port = value
			delete(meta, key)
			break
		case PATH:
			path = value
			delete(meta, key)
			break
		case PARAMS:
			params = value
			delete(meta, key)
			break
		}
	}

	if "" == method {
		return nil, errors.New("Can not build http request, \"method\" is missing.")
	}

	if "" == protocol && !strings.HasPrefix(host, "http") {
		protocol = "http://"
	}

	if "" == host {
		return nil, errors.New("Can not build http request, \"host\" is missing.")
	}

	if "" == port && !strings.Contains(strings.Replace(host, "://", "", 0), ":") {
		path = ":80"
	}

	if "" == path {
		path = "/"
	}

	return &Request{
		headers:     meta,
		method:      method,
		protocol:    protocol,
		host:        host,
		port:        port,
		path:        path,
		params:      params,
		body:        req.GetPayload(),
		baseRequest: req,
	}, nil
}

func (self *Request) GetHeaders() map[string]string {
	return self.headers
}

func (self *Request) GetMethod() string {
	return self.method
}

func (self *Request) GetProtocol() string {
	return self.protocol
}

func (self *Request) GetHost() string {
	return self.host
}

func (self *Request) GetPort() string {
	return self.port
}

func (self *Request) GetPath() string {
	return self.path
}

func (self *Request) GetParams() string {
	return self.params
}

func (self *Request) GetBody() []byte {
	return self.body
}

func (self *Request) toBaseRequest() communication.RequestAware {
	return self.baseRequest
}
