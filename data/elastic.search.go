package data

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/settings"
)

var ElasticSearchClient *elasticsearch.Client

func init() {
	config := elasticsearch.Config{
		Addresses: []string{
			settings.Configuration.GetString("es1.connectionstring"),
		},
	}

	var err error
	ElasticSearchClient, err = elasticsearch.NewClient(config)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := ElasticSearchClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}

func GetInfo() {
	log.Printf("Client: %s", elasticsearch.Version)
}
