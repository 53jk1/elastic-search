package app

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"strings"
)

// App is the main application
type App struct {
	es *elasticsearch.Client
}

// NewApp creates a new App
func NewApp() *App {
	return &App{}
}

// Run runs the application
func (a *App) Run() error {
	// Create an Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	a = &App{es: es}

	// Create an index
	ctx := context.Background()
	res, err := a.es.Indices.Create(
		"myindex",
		a.es.Indices.Create.WithBody(strings.NewReader(`{"settings": {"number_of_shards": 1,"number_of_replicas": 0}}`)),
		a.es.Indices.Create.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	fmt.Println(res)

	// Index a document
	res, err = a.es.Index(
		"myindex",
		strings.NewReader(`{"title" : "Test"}`),
		a.es.Index.WithDocumentID("1"),
		a.es.Index.WithRefresh("true"),
	)
	if err != nil {
		return err
	}
	fmt.Println(res)

	// Get a document
	res, err = a.es.Get(
		"myindex",
		"1",
		a.es.Get.WithPretty(),
	)
	if err != nil {
		return err
	}
	fmt.Println(res)

	// Delete an index
	res, err = a.es.Indices.Delete([]string{"myindex"})
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
