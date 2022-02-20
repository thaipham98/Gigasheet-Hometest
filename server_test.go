package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ipsumHandler(t *testing.T) {
	var jsonStr = []byte(`{"ips": ["94.142.241.194","192.168.1.1","159.65.180.64"]}`)
	request, err := http.NewRequest("GET", "/count_ips_in_ipsum", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(ipsumHandler)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "2\n"
	if response.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}
