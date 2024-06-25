package main

import (
	"fmt"
	"log"

	"github.com/tasnimzotder/go-junitxml"
)

var (
	xmlGo    = "cmd/test/junit.xml"
	xmlRuby  = "cmd/test/report.xml"
	xmlRuby2 = "cmd/test/report2.xml"
)

func main() {
	junit := junitxml.NewJUnitXML()

	suites, err := junit.ParseFile(xmlRuby)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", suites.Totals)

	// mergedSuites, err := junitxml.Merge(suites)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v\n", mergedSuites.Totals)

	mergedSuites, err := junit.MergeFiles(xmlRuby, xmlGo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", mergedSuites.Totals)

	opts := &junitxml.MergeOptions{
		SingleTestSuite: false,
	}

	mergedSuites, err = junit.MergeFilesWithOpts(opts, xmlRuby, xmlGo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", mergedSuites.Totals)
}
