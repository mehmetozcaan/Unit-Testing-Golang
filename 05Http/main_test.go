package testhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {

	request := httptest.NewRequest("GET", "/sayhello", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, request)

	// Check Response Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Http Request Failed! Response Code: %v", status)
	}

	// Check Response Body
	response := rr.Body.String()
	if response != "hello" {
		t.Errorf("Unexpected Response Message! : %s", response)
	}

}
