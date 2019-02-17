package main

import (
	"fmt"
	"github.com/phouse512/coda-go"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
)

type Transport struct {
	defaultTransport http.RoundTripper
	token            string
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", "Bearer", t.token))
	return t.defaultTransport.RoundTrip(req)
}

func main() {
	fmt.Println("Starting coda client api test.")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Unable to load config.")
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &Transport{
			defaultTransport: http.DefaultTransport,
			token:            viper.GetString("coda_api_key"),
		},
	}

	u, err := url.Parse("https://coda.io/apis/v1beta1")
	if err != nil {
		log.Print("Unable to parse url")
		log.Fatal(err)
	}

	codaClient := coda.Client{
		UserAgent:  "TestAgent",
		HttpClient: client,
		BaseURL:    u,
	}

	result, err := codaClient.ListDocs()
	if err != nil {
		log.Print("Unable to list documents")
		log.Fatal(err)
	}

	for _, doc := range result {
		log.Print("Doc ID: ", doc.Id)
		log.Print("Doc Name: ", doc.Name)

		document, err := codaClient.GetDoc(doc.Id)
		if err != nil {
			log.Print("Unable to get document")
			log.Fatal(err)
		}

		log.Print("DocumentLink: ", document.Document.BrowserLink)
	}
}
