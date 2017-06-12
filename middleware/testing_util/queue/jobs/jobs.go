package jobs

import (
	"gyg/middleware/logger"
	"gyg/middleware/processing/map_reduce"
	jobsLib "gyg/middleware/queue/jobs"
	"time"
)

func dummyMapper(input interface{}, output chan interface{}, loggerFunc logger.LoggerFunc, errorsChannel chan error) {
}

func GetDummyJob() *jobsLib.Job {
	mappers := []map_reduce.MapperFunc{dummyMapper}
	payload := jobsLib.NewPayload(map[string]interface{}{})
	output := make(chan interface{}, 1)
	return jobsLib.NewJob("Test-Job", 0, time.Now(), payload, mappers, output)
}
