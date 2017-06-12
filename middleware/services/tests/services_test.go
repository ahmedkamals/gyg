package tests

import (
	"errors"
	"gyg/middleware/config"
	"gyg/middleware/services"
	testingUtil "gyg/middleware/testing_util"
	"reflect"
	"testing"
)

func TestShouldLoadConfigCorrectly(t *testing.T) {
	testCases := []*testingUtil.TestCase{
		&testingUtil.TestCase{
			Id:       "Should return error.",
			Input:    []byte(``),
			Expected: errors.New("unexpected end of JSON input"),
		},
		&testingUtil.TestCase{
			Id:       "Should parse config correctly.",
			Input:    testingUtil.GetConfigString(),
			Expected: testingUtil.GetConfig(),
		},
	}

	for _, testCase := range testCases {

		input := testCase.Input.([]byte)
		serviceLocator := services.NewLocator()
		got, err := serviceLocator.LoadConfig(input)

		if err != nil && testCase.Expected.(error).Error() != err.Error() {
			t.Error(testingUtil.Format(testCase.Id, testCase.Expected, err))
		}
		if err == nil && !reflect.DeepEqual(testCase.Expected.(*config.Config), got) {
			t.Error(testingUtil.Format(testCase.Id, testCase.Expected, err))
		}
	}
}
