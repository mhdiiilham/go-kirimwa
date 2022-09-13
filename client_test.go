package gokirimwa_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	gokirimwa "github.com/mhdiiilham/go-kirimwa"
	"github.com/stretchr/testify/assert"
)

var API_KEY = faker.CCNumber()

func TestNewKirimWA(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := gokirimwa.NewKirimWA(API_KEY)

		assertion := assert.New(t)
		assertion.NotNil(c)
	})
}
