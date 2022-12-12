package main

import (
	"github.com/53jk1/elastic-search/app"
)

func main() {
	x := app.NewApp()
	if err := x.Run(); err != nil {
		panic(err)
	}
}
