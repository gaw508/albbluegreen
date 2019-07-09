package albbluegreen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvertStatus(t *testing.T) {
	testCases := map[string]struct {
		status         BlueGreenStatus
		expectedStatus BlueGreenStatus
	}{
		"blue -> green": {
			status: "blue",
			expectedStatus: "green",
		},
		"green -> blue": {
			status: "green",
			expectedStatus: "blue",
		},
		"defaults to blue": {
			status: "notacolour",
			expectedStatus: "blue",
		},
		"defaults to blue (empty)": {
			status: "",
			expectedStatus: "blue",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			status := InvertStatus(tc.status)
			assert.Equal(t, tc.expectedStatus, status)
		})
	}
}
