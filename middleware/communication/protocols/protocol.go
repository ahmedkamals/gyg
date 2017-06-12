package protocols

import (
	"gyg/middleware/communication"
)

type Protocol interface {
	Send(*communication.Request) (*communication.Response, error)
}
