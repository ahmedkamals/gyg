package testing_util

import (
	"fmt"
	"github.com/kr/pretty"
	"net/http"
	"net/http/httptest"
	"os"
	"gyg/middleware/config"
	"gyg/middleware/logger"
	"gyg/middleware/processing/geolocation"
	"time"
)

func GetMockServer(handler func(http.ResponseWriter, *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}

func GetLoggerFunc() func(logger.Level, string) {
	return func(level logger.Level, message string) {

	}
}

func Format(caseId string, expected, got interface{}) []string {
	colorable := logger.NewColorable(os.Stderr)

	caseId = colorable.Wrap(fmt.Sprintf("\nCase:\t  %q", caseId), logger.CYAN)
	pretty.Println(caseId)

	return pretty.Diff(expected, got)
}

func GetConfigString() []byte {
	return []byte(`{
  "active_version": "inrix-v1",
  "results_file_name": "data.merged.json",
  "processing": {
    "max_workers": 19,
    "max_queued_items": 500
  },
  "locations": {
    "min_weight": 3,
    "scan_range": {
      "range": {
        "start": {
          "latitude": 71.183768,
          "longitude": 40.100098
        },
        "end": {
          "latitude": 35.994535,
          "longitude": -10.557861
        }
      },
      "lat_step": 0.5,
      "long_step": 0.5
    }
  },
  "phrase_app": {
    "token": "Phrase that bitch.",
    "project_id": "Phrase it so much.",
    "enabled": false
  },
  "inrix-v2": {
    "authentication" : {
      "method": "GET",
      "protocol": "https://",
      "host": "uas-api.beta.inrix.com/v1/AppToken",
      "port": ":80",
      "path": "/",
      "params": {
	"appId": "APP_ID",
	"appKey": "APP_KEY"
      },
       "retries_limit": 3,
       "retry_waiting_duration": 5000
    },
    "locations" : {
        "method": "GET",
        "protocol": "https://",
        "host": "parking-api.inrix.com",
        "port": ":80",
        "path": "/lots",
        "retries_limit": 3,
        "retry_waiting_duration": 5000
    }
  },
  "inrix-v1": {
    "locations" : {
      "method": "GET",
      "protocol": "http://",
      "host": "eu.api.inrix.com",
      "port": ":80",
      "path": "/Traffic/Inrix.ashx",
      "params": {
        "token": "TOKEN",
        "action": "GetParkingInfoInBox",
        "locale": "en-GB",
        "compress": "true"
      },
       "retries_limit": 3,
       "retry_waiting_duration": 5000
    }
  }
}`)
}

func GetConfig() *config.Config {
	return &config.Config{
		ActiveVersion:   "inrix-v1",
		ResultsFileName: "data.merged.json",
		Processing: config.Processing{
			MaxWorkers:     19,
			MaxQueuedItems: 500,
		},
		Locations: config.Location{
			MinWeight: 3,
			ScanRange: config.ScanRange{
				Range: geolocation.BoundingBox{
					Start: geolocation.Coordinates{
						Latitude:  71.183768,
						Longitude: 40.100098,
					},
					End: geolocation.Coordinates{
						Latitude:  35.994535,
						Longitude: -10.557861,
					},
				},
				LatStep:  0.5,
				LongStep: 0.5,
			},
		},
		PhraseApp: config.PhraseApp{
			Token:     "Phrase that bitch.",
			ProjectId: "Phrase it so much.",
			Enabled:   false,
		},
		Backend: map[string]*config.APIEndPoint{
			"locations": &config.APIEndPoint{
				Method:   "GET",
				Protocol: "http://",
				Host:     "eu.api.inrix.com",
				Port:     ":80",
				Path:     "/Traffic/Inrix.ashx",
				Params: map[string]string{
					"token":    "TOKEN",
					"action":   "GetParkingInfoInBox",
					"locale":   "en-GB",
					"compress": "true",
				},
				RetriesLimit:         3,
				RetryWaitingDuration: time.Duration(5000),
			},
		},
		InrixV2: map[string]*config.APIEndPoint{
			"authentication": &config.APIEndPoint{
				Method:   "GET",
				Protocol: "https://",
				Host:     "uas-api.beta.inrix.com/v1/AppToken",
				Port:     ":80",
				Path:     "/",
				Params: map[string]string{
					"appId":  "APP_ID",
					"appKey": "APP_KEY",
				},
				RetriesLimit:         3,
				RetryWaitingDuration: time.Duration(5000),
			},
			"locations": &config.APIEndPoint{
				Method:   "GET",
				Protocol: "https://",
				Host:     "parking-api.inrix.com",
				Port:     ":80",
				Path:     "/lots",
				//Params: nil,
				RetriesLimit:         3,
				RetryWaitingDuration: time.Duration(5000),
			},
		},
	}
}
