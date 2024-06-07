package main

import (
	"bytes"
	"goldWatcher/repository"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.HTTPClient = client
	testApp.MainWindow = a.NewWindow("")
	testApp.DB = repository.NewTestReposotory()

	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1717507701036,
	"tsj": 1717507696086,
	"date": "Jun 4th 2024, 09:28:16 am NY",
	"items": [
	  {
		"curr": "USD",
		"xauPrice": 2337.895,
		"xagPrice": 29.9579,
		"chgXau": -12.765,
		"chgXag": -0.7771,
		"pcXau": -0.543,
		"pcXag": -2.5284,
		"xauClose": 2350.66,
		"xagClose": 30.735
	  }
	]
  }
`

type RoundTricFunc func(req *http.Request) *http.Response

func (f RoundTricFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTricFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
