# nucredit

名大の修得科目jsonファイルを扱うライブラリ

## install

`$ go get "github.com/nu50218/nucredit"`

## sample code

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nu50218/nucredit"
)

func main() {
	f, err := os.Open("subjects.json")
	if err != nil {
		log.Fatal(err)
	}

	s, err := nucredit.FromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total GPA:", s.GPA())

	s2019 := s.Filter(func(sb *nucredit.Subject) bool {
		return sb.Year == 2019
	})
	fmt.Println("2019 GPA:", s2019.GPA())
}
```
