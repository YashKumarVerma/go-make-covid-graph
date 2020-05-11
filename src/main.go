package main

import (
	"github.com/YashKumarVerma/go-make-covid-graph/src/packages/data"
	"github.com/YashKumarVerma/go-make-covid-graph/src/packages/grapher"
)

func main() {
	latestData := data.Fetch()
	grapher.Generate(latestData)
}
