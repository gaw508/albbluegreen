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

func TestToggleHandler_Success(t *testing.T) {
	testCases := map[string]struct {
		newStatus              albbluegreen.BlueGreenStatus
		serviceError           error
		expectedCliOutput      string
		expectedError          error
		expectedServiceInvoked bool
	}{
		"to green": {
			newStatus:              albbluegreen.Green,
			serviceError:           nil,
			expectedCliOutput:      "Successfully toggled to: green\n",
			expectedError:          nil,
			expectedServiceInvoked: true,
		},
		"service error": {
			newStatus:              "",
			serviceError:           assert.AnError,
			expectedCliOutput:      "",
			expectedError:          cli.NewExitError("Failed to toggle status: assert.AnError general error for testing", 1),
			expectedServiceInvoked: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockService := &mock.BlueGreenService{
				ToggleFn: func() (status albbluegreen.BlueGreenStatus, err error) {
					return tc.newStatus, tc.serviceError
				},
			}
			logBuffer := &bytes.Buffer{}
			logger := log.New(logBuffer, "", 0)

			statusHandler := ToggleHandler(logger, mockService)
			err := statusHandler(&cli.Context{})

			assert.Equal(t, tc.expectedServiceInvoked, mockService.ToggleInvoked)
			assert.Equal(t, tc.expectedCliOutput, logBuffer.String())
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
