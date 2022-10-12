package gokirimwa_test

import (
	"errors"
	"testing"

	gokirimwa "github.com/mhdiiilham/go-kirimwa"
	"github.com/stretchr/testify/assert"
)

func TestErrorResponseError(t *testing.T) {
	t.Run("return string", func(t *testing.T) {
		errResp := gokirimwa.ErrorResponse{Message: "Error: Device not found."}
		expected := "Error: Device not found."

		assert.Equal(t, expected, errResp.Error())
	})
}

func TestIsKirimWAError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "given error is gokirimwa's error",
			err:      gokirimwa.ErrorResponse{Message: "Error: Device not found."},
			expected: true,
		},
		{
			name:     "given error is not gokirimwa's error",
			err:      errors.New("error"),
			expected: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			actual := gokirimwa.IsKirimWAError(tt.err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
