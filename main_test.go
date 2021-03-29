package main

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestStatusType(t *testing.T) {
	client := resty.New()
	resp, _ := client.R().Get("http://localhost:8080/orders")
	assert.Equal(t, 200, resp.StatusCode())
}
