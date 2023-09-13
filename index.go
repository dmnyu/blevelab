package main

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
)

func CreateIndex(indexName string) error {
	mapping := bleve.NewIndexMapping()
	pointMapping := bleve.NewDocumentMapping()
	xFieldMapping := bleve.NewNumericFieldMapping()
	yFieldMapping := bleve.NewNumericFieldMapping()
	nameMapping := bleve.NewTextFieldMapping()
	nameMapping.Analyzer = keyword.Name
	pointMapping.AddFieldMappingsAt("X", xFieldMapping)
	pointMapping.AddFieldMappingsAt("Y", yFieldMapping)
	pointMapping.AddFieldMappingsAt("Name", nameMapping)
	mapping.AddDocumentMapping("point", pointMapping)
	index, err := bleve.New(indexName, mapping)
	if err != nil {
		return err
	}
	fmt.Println("created new index", index.Name())

	if err := index.Close(); err != nil {
		return err
	}

	return nil
}

func PopulateIndex(indexName string) error {

	index, err := bleve.Open(indexName)
	if err != nil {
		return err
	}

	point := Point{1, 2, "First Point"}
	point2 := Point{4, 2, "Second Point"}
	point3 := Point{9, 9, "Third Point"}
	if err := index.Index("Point1", point); err != nil {
		return err
	}

	if err := index.Index("Point2", point2); err != nil {
		return err
	}

	if err := index.Index("Point3", point3); err != nil {
		return err
	}

	if err := index.Close(); err != nil {
		return err
	}
	return nil
}
