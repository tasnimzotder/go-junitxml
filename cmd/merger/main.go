package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tasnimzotder/go-junitxml"
)

type stringArray []string

func (s *stringArray) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var inputFiles stringArray
	flag.Var(&inputFiles, "input", "input files to merge")
	output := flag.String("output", "merged-junit.xml", "output file")
	flag.Parse()

	log.Println("input files:", inputFiles)
	log.Println("output file:", *output)

	if len(inputFiles) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	junit := junitxml.NewJUnitXML()

	merged, err := junit.MergeFiles(inputFiles...)
	if err != nil {
		log.Fatalf("failed to merge files: %v", err)
	}

	err = junit.WriteToFile(merged, *output)
	if err != nil {
		log.Fatalf("failed to write to file: %v", err)
	}
}
