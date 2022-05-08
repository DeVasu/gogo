package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/unrolled/render"
)

const (
	fakeMatchLocationResult = "/matches/5a003b78-409e-4452-b456-a6f0dcee05bd"
)

var (
	formatter = render.New(render.Options{ 
		IndentJSON: true,
	})
)

func TestCreateMatch(t *testing.T) { 
	client := &http.Client{}
	server := httptest.NewServer( http.HandlerFunc(createMatchHandler(formatter)))
	defer server.Close()
	body := []byte("{\n \"gridsize\": 19,\n \"players\": [\n {\n \"color\": \"white\",\n \"name\": \"bob\"\n },\n {\n \"color\": \"black\",\n \"name\": \"alfred\"\n }\n ]\n}")

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createMatchHandler: %v",err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req) 
	if err != nil {
		t.Errorf("Error in POST to createMatchHandler: %v", err)
	}
	defer res.Body.Close()
	payload, err := ioutil.ReadAll(res.Body) 
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s",res.Status)
	}
	fmt.Printf("Payload: %s", string(payload))

<<<<<<< HEAD
	if _, ok := res.Header["Location"]; !ok {
<<<<<<< HEAD
		t.Error("Location Header not set")
=======
		 t.Error("Location header is not set")
>>>>>>> a37e028ffb3fe6f64396522331f39a45f585ae5e
=======
	loc, headerOk := res.Header["Location"]; 
	if !headerOk {
		 t.Error("Location header is not set")
	} else {
		if !strings.Contains(loc[0], "/matches/") {
			t.Errorf("Location header should contain '/matches/'")
		}
		if len(loc[0]) != len(fakeMatchLocationResult) {
			t.Errorf("Location values does not contain guid of new match")
		}
>>>>>>> 5e9fdf00c00b9b5e5f0567ede6afcb579fec141c
	}
}