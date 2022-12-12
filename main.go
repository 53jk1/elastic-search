package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"strings"
)

func main() {
	// Create an Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
	}

	// Create an index
	ctx := context.Background()
	res, err := es.Indices.Create(
		"myindex",
		es.Indices.Create.WithBody(strings.NewReader(`{
			"settings": {
				"number_of_shards": 1,
				"number_of_replicas": 0
			}
		}`)),
		es.Indices.Create.WithContext(ctx),
	)
	if err != nil {
		fmt.Println(err)
	}

	// Index a document
	res, err = es.Index(
		"myindex",
		strings.NewReader(`{"title" : "Test"}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		fmt.Println(err)
	}
	// Get a document
	res, err = es.Get(
		"myindex",
		"1",
		es.Get.WithPretty(),
	)
	if err != nil {
		fmt.Println(err)
	}

	// Delete an index
	res, err = es.Indices.Delete([]string{"myindex"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
