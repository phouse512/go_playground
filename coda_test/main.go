package main

import (
	"fmt"
	"github.com/phouse512/coda-go"
	"log"
	"net/http"
	"net/url"
)

type Transport struct {
	defaultTransport http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", "Bearer", "fake-client"))
	return t.defaultTransport.RoundTrip(req)
}

func main() {
	fmt.Println("vim-go")

	client := &http.Client{
		Transport: &Transport{
			defaultTransport: http.DefaultTransport,
		},
	}

	u, err := url.Parse("https://coda.io/apis/v1beta1")
	if err != nil {
		log.Fatal(err)
	}

	codaClient := coda.Client{
		UserAgent:  "TestAgent",
		HttpClient: client,
		BaseURL:    u,
	}
}
