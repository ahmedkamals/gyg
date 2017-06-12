package communication

type (
	ResponseAware interface {
		GetPayload() string
		GetMeta() map[string]string
		IsSuccessful() bool
	}

	Response struct {
		payload      string
		meta         map[string]string
		isSuccessful bool
		ResponseAware
	}
)

func NewResponse(payload string, meta map[string]string, isSuccessful bool) *Response {
	return &Response{
		payload:      payload,
		meta:         meta,
		isSuccessful: isSuccessful,
	}
}

func (resp *Response) GetPayload() string {
	return resp.payload
}

func (resp *Response) GetMeta() map[string]string {
	return resp.meta
}

func (resp *Response) IsSuccessful() bool {
	return resp.isSuccessful
}
