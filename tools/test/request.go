package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

type IMockRequest interface {
	Get(path string, body interface{}) *http.Request
}

type mockRequest struct {
}

func MockRequest() IMockRequest {
	return &mockRequest{}
}

func (m *mockRequest) Get(path string, body interface{}) *http.Request {

	return httptest.NewRequest("GET", path, strings.NewReader(body.(string)))
}
