package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestResponse(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    SayHello(response, request)

    t.Run("Return Hello, world!\n", func(t *testing.T){
        got := response.Body.String()
        expected := "Hello, world!\n"

        if got != expected {
            t.Errorf("Got %q, expected %q", got, expected)
        }
    })
}
