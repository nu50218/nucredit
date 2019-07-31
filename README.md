# nucredit

名大の修得科目jsonファイルを扱うライブラリ

## sample code

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nu50218/nucredit"
)

func main() {
	bytes, err := ioutil.ReadFile("credit.json")
	if err != nil {
		log.Fatalln(err)
	}
	subjects := nucredit.Subjects{}
	err = json.Unmarshal(bytes, &subjects)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Total GPA:", subjects.CalcGradePointAverage())
	fmt.Println("2019 GPA:", subjects.
		Filter(func(s *nucredit.Subject) bool { return s.Year == 2019 }).
		CalcGradePointAverage(),
	)
}
```
