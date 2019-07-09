package cli

import (
	"bytes"
	"flag"
	"github.com/gaw508/albbluegreen"
	"github.com/gaw508/albbluegreen/mock"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"log"
	"testing"
)

func TestSetStatusHandler_Success(t *testing.T) {
	testCases := map[string]struct {
		newStatus              albbluegreen.BlueGreenStatus
		serviceError           error
		expectedCliOutput      string
		expectedError          error
		expectedServiceInvoked bool
	}{
		"blue": {
			newStatus:              "blue",
			serviceError:           nil,
			expectedCliOutput:      "Successfully set service to: blue\n",
			expectedError:          nil,
			expectedServiceInvoked: true,
		},
		"empty": {
			newStatus:              "",
			serviceError:           nil,
			expectedCliOutput:      "",
			expectedError:          cli.NewExitError("status is required", 1),
			expectedServiceInvoked: false,
		},
		"service error": {
			newStatus:              "blue",
			serviceError:           assert.AnError,
			expectedCliOutput:      "",
			expectedError:          cli.NewExitError("Failed to set status: assert.AnError general error for testing", 1),
			expectedServiceInvoked: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockService := &mock.BlueGreenService{
				SetStatusFn: func(status albbluegreen.BlueGreenStatus) error {
					assert.Equal(t, tc.newStatus, status)
					return tc.serviceError
				},
			}
			logBuffer := &bytes.Buffer{}
			logger := log.New(logBuffer, "", 0)

			fs := &flag.FlagSet{}
			fs.String("status", "", "")
			ctx := cli.NewContext(nil, fs, nil)
			_ = ctx.Set("status", string(tc.newStatus))

			statusHandler := SetStatusHandler(logger, mockService)
			err := statusHandler(ctx)

			assert.Equal(t, tc.expectedServiceInvoked, mockService.SetStatusInvoked)
			assert.Equal(t, tc.expectedCliOutput, logBuffer.String())
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
