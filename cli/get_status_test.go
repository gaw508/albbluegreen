package cli

import (
	"bytes"
	"github.com/gaw508/albbluegreen"
	"github.com/gaw508/albbluegreen/mock"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"log"
	"testing"
)

func TestGetStatusHandler_Success(t *testing.T) {
	testCases := map[string]struct {
		currentStatus          albbluegreen.BlueGreenStatus
		serviceError           error
		expectedCliOutput      string
		expectedError          error
		expectedServiceInvoked bool
	}{
		"blue": {
			currentStatus:          albbluegreen.Blue,
			serviceError:           nil,
			expectedCliOutput:      "Current active service: blue\n",
			expectedError:          nil,
			expectedServiceInvoked: true,
		},
		"green": {
			currentStatus:          albbluegreen.Green,
			serviceError:           nil,
			expectedCliOutput:      "Current active service: green\n",
			expectedError:          nil,
			expectedServiceInvoked: true,
		},
		"service error": {
			currentStatus:          "",
			serviceError:           assert.AnError,
			expectedCliOutput:      "",
			expectedError:          cli.NewExitError("Failed to get status: assert.AnError general error for testing", 1),
			expectedServiceInvoked: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockService := &mock.BlueGreenService{
				StatusFn: func() (status albbluegreen.BlueGreenStatus, err error) {
					return tc.currentStatus, tc.serviceError
				},
			}
			logBuffer := &bytes.Buffer{}
			logger := log.New(logBuffer, "", 0)

			statusHandler := GetStatusHandler(logger, mockService)
			err := statusHandler(&cli.Context{})

			assert.Equal(t, tc.expectedServiceInvoked, mockService.StatusInvoked)
			assert.Equal(t, tc.expectedCliOutput, logBuffer.String())
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
