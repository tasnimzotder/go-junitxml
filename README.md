# JUnitXML

JUnitXML is a Go package for parsing, validating, writing, and merging JUnit XML files.

## Usage

### Installation

```bash
go get github.com/tasnimzotder/go-junit/pkg
```

### Parsing

```go
package main

import (
	"fmt"
	"github.com/tasnimzotder/go-junitxml"
)

func main() {
	junit := junitxml.NewJUnitXML()

	suite, err := junit.ParseFile("junit.xml")
	if err != nil {
		panic(err)
	}

	fmt.Println(suite)
}
```

### Merging

```go
package main

import (
	"github.com/tasnimzotder/go-junitxml"
)

func main() {
	junit := junitxml.NewJUnitXML()

	suite1, _ := junit.ParseFile("juint1.xml")
	suite2, _ := junit.ParseFile("juint2.xml")

	merged, _ := junit.Merge(suite1, suite2)

	_ = junit.WriteToFile(merged, "merged.xml")
}
```
