package gokirimwa

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendQueries(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queries := map[string]string{"limit": "100", "page": "2"}
		url := "http://localhost:8000/users"
		expectedRawQuery := "limit=100&page=2"

		req, _ := http.NewRequest(http.MethodGet, url, nil)
		c := NewKirimWA("fake_api_key")
		actual := c.appendQueries(req, queries)

		assert.Equal(t, expectedRawQuery, actual)
	})

	t.Run("queries is empty", func(t *testing.T) {
		queries := map[string]string{}
		url := "http://localhost:8000/users"
		expectedRawQuery := ""

		req, _ := http.NewRequest(http.MethodGet, url, nil)
		c := NewKirimWA("fake_api_key")
		actual := c.appendQueries(req, queries)

		assert.Equal(t, expectedRawQuery, actual)
	})

	t.Run("queries is nil", func(t *testing.T) {
		queries := map[string]string{}
		url := "http://localhost:8000/users"
		expectedRawQuery := ""

		req, _ := http.NewRequest(http.MethodGet, url, nil)
		c := NewKirimWA("fake_api_key")
		actual := c.appendQueries(req, queries)

		assert.Equal(t, expectedRawQuery, actual)
	})
}
