# Go Parser
Simple parse JSON/CSV request

### Install
`go get github.com/marioidival/go_parser`

### Usage
main.go:


```go
package main

import (
    "net/http",
    "github.com/marioidival/go_parser"
    "fmt"
)

//  jsonContent = "http://www.mocky.io/v2/556b0ba673eedce302329da9"
//  jsonExtraContent = "http://www.mocky.io/v2/556b214973eedc9503329dae

func main() {
    url := os.Args[1]
    
    resp, _ := http.Get(url)
    
    pars := &parser.Parser{Resp: resp}
    pars.GetContent() // Verify if content is application/json or text/csv
    mapped := pars.ParseBody() // return a map with fields and values coming from request
    fmt.Printf("%v\n", mapped)
}
```

```
go run main.go http://www.mocky.io/v2/556b0ba673eedce302329da9
[map[sex:M name:mario age:23 email:marioidival@gmail.com]
```
