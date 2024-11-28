package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		contentType    string
		body           string
		expectedStatus int
		expectedBody   []string
	}{
		{
			name:           "Successful POST request",
			method:         "POST",
			contentType:    "application/x-www-form-urlencoded",
			body:           "name=John&address=123+Main+St",
			expectedStatus: http.StatusOK,
			expectedBody:   []string{"POST request successful", "Name = John", "Address = 123 Main St"},
		},
		{
			name:           "ParseForm error",
			method:         "POST",
			contentType:    "application/",
			body:           "name=John&address=123%Main%St",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []string{"ParseForm() err"},
		},
		{
			name:           "Unsupported method",
			method:         "GET",
			contentType:    "",
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   []string{"method is not supported"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/form", strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", tt.contentType)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(formHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			for _, expected := range tt.expectedBody {
				assert.Contains(t, rr.Body.String(), expected)
			}
		})
	}
}

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful GET request",
			url:            "/hello",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   "hello!",
		},
		{
			name:           "Invalid URL",
			url:            "/invalid",
			method:         "GET",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 not found",
		},
		{
			name:           "Unsupported method",
			url:            "/hello",
			method:         "POST",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "method is not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(helloHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Contains(t, rr.Body.String(), tt.expectedBody)
		})
	}
}

func TestServeForm(t *testing.T) {
	req, err := http.NewRequest("GET", "/form", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveForm)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "<form")
}

func TestSetupServer(t *testing.T) {
	mux := setupServer()

	tests := []struct {
		method      string
		path        string
		contentType string
		body        string
		want        int
	}{
		{"GET", "/", "", "", http.StatusOK},
		{"POST", "/form", "application/x-www-form-urlencoded", "name=John&address=123+Main+St", http.StatusOK},
		{"POST", "/api/v1/form", "application/x-www-form-urlencoded", "name=John&address=123+Main+St", http.StatusOK},
		{"GET", "/api/v1/hello", "", "", http.StatusOK},
	}

	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, tt.path, nil)
		if tt.body != "" {
			req, err = http.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
		}

		if err != nil {
			t.Fatal(err)
		}

		if tt.contentType != "" {
			req.Header.Set("Content-Type", tt.contentType)
		}

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.want {
			t.Errorf("handler %s %s returned wrong status code: got %v want %v",
				tt.method, tt.path, status, tt.want)
		}
	}
}
