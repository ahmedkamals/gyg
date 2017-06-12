package config

import (
	"time"
)

type (
	Server struct {
		Host, Port, Path string
	}

	APIEndPoint struct {
		Method, Protocol, Host, Port, Path string
		Params                             map[string]string
		// The number of retries allowed for the request.
		RetriesLimit uint `json:"retries_limit"`
		// The waiting duration in milli-seconds before the next retry.
		RetryWaitingDuration time.Duration `json:"retry_waiting_duration"`
	}

	Config struct {
		Server          Server 			`json:"server"`
		Backend         map[string]APIEndPoint `json:"backend"`
		RatingFile	string			`json:"rating_file"`
	}
)

var (
	Configuration *Config
)
