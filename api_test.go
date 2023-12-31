package main

import (
	"log"
	"net/http"
	"testing"
)

func TestGetRoute(t *testing.T) {
	url := "http://localhost:8000/api/v1/restaurant"
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	log.Println("Status:", resp.Status)
	log.Println("Body:", resp.Body)
}
