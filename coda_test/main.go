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
	}

	testDocId := "wPeV8NJTam"
	document, err := codaClient.GetDoc(testDocId)
	if err != nil {
		log.Print("Unable to get test document.")
		log.Fatal(err)
	}
	log.Print(document)

	sectionsPayload := coda.ListSectionsPayload{Limit: 2, PageToken: "2"}
	sectionsResponse, err := codaClient.ListSections(testDocId, sectionsPayload)
	if err != nil {
		log.Print("unable to list sections.")
		log.Fatal(err)
	}

	sectionResponse, err := codaClient.GetSection(testDocId, sectionsResponse.Sections[0].Id)
	if err != nil {
		log.Print("Unable to get section.")
		log.Fatal(err)
	}
	log.Print(sectionResponse)

	foldersResponse, err := codaClient.ListFolders(testDocId, coda.PaginationPayload{})
	if err != nil {
		log.Print("Unable to list folders.")
		log.Fatal(err)
	}
	log.Print(foldersResponse)

	tablesResp, err := codaClient.ListTables(testDocId, coda.PaginationPayload{})
	if err != nil {
		log.Print("Unable to list tables.")
		log.Fatal(err)
	}

	tableDetail, err := codaClient.GetTable(testDocId, tablesResp.Tables[0].Id)
	log.Print(tableDetail)

	viewsResponse, err := codaClient.ListViews(testDocId, coda.PaginationPayload{})
	if err != nil {
		log.Print("unable to list views.")
		log.Fatal(err)
	}

	log.Print(viewsResponse)

	colsResp, err := codaClient.ListColumns(testDocId, tableDetail.Table.Id, coda.PaginationPayload{})
	log.Print(colsResp)
	for _, col := range colsResp.Columns {
		log.Print(fmt.Sprintf("Column Id: %s and Name: %s", col.Id, col.Name))
	}
	/*	document, err := codaClient.GetDoc(doc.Id)
		if err != nil {
			log.Print("Unable to get document")
			log.Fatal(err)
		}

		sectionsResp, err := codaClient.ListSections(doc.Id)
		if err != nil {
			log.Print("Unable to fetch responses.")
			log.Fatal(err)
		}

		log.Print("DocumentLink: ", document.Document.BrowserLink)
		log.Print("Sections: ", sectionsResp.PaginationResponse.NextPageLink)
	}*/

	//	var createPayload = coda.CreateDocPayload{
	//		Title: "FAKE DOC v2",
	//	}
	//	createResult, err := codaClient.CreateDoc(createPayload)
	//	log.Print(createResult)
	//
	log.Print("Rows tests.")
	rowsResp, err := codaClient.ListTableRows(testDocId, tableDetail.Table.Id, coda.ListRowsParameters{})
	log.Print(rowsResp.Rows[0])

	insertRowsParams := coda.InsertRowsParameters{
		Rows: [
			{
				Cells: [
					{
						Column: "c-XZAGElgNkD",
						Value: 123,
					},
				]
			}
		],
	}
}
