package testing_util

import (
	"fmt"
	"github.com/kr/pretty"
	"net/http"
	"net/http/httptest"
	"os"
	"gyg/middleware/config"
	"gyg/middleware/logger"
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

