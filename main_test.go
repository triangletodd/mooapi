package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Moo"
	actual, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if string(actual) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}

}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Response{Message: "Hello, World!"}
	var actual Response
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestGreetHandler(t *testing.T) {
	tests := []struct {
		name     string
		param    string
		expected string
	}{
		{"NoName", "", "Hello, Guest!"},
		{"WithName", "John", "Hello, John!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/greet"
			if tt.param != "" {
				url += "?name=" + tt.param
			}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(greetHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			expected := Response{Message: tt.expected}
			var actual Response
			if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
				t.Fatal(err)
			}

			if actual != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					actual, expected)
			}
		})
	}
}
