package main

import (
	"flag"
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

var (
	query  = flag.String("query", "", "")
	create = flag.Bool("create", false, "")
	index  bleve.Index
)

type Point struct {
	X    int
	Y    int
	Name string
}

func main() {
	flag.Parse()
	indexName := "point-index"

	if *create {
		if err := CreateIndex(indexName); err != nil {
			panic(err)
		}

		fmt.Println("populating index")
		if err := PopulateIndex(indexName); err != nil {
			panic(err)
		}
	}
	//Open the index
	fmt.Println("Opening index")
	index, err := bleve.Open(indexName)
	if err != nil {
		panic(err)
	}

	//create a query
	q := bleve.NewFuzzyQuery(*query)
	search := bleve.NewSearchRequest(q)
	search.Fields = []string{"*"}

	//perform a search
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}

	//itterate through results
	hits := searchResults.Hits
	if len(hits) > 0 {
		fmt.Println("Results\n=======")
		for _, hit := range hits {
			fmt.Println(GetPoint(hit.Fields))
		}
	}
}

// marshal a index hit to a point
func GetPoint(hitMap map[string]interface{}) Point {
	name := hitMap["Name"].(string)
	x := int(hitMap["X"].(float64))
	y := int(hitMap["Y"].(float64))

	return Point{x, y, name}
}
