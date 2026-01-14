## Usage

### Install

```bash
go get github.com/ayushkr12/xurl@latest
```

### Import & Use

```go
package main

import (
	"fmt"
	"log"

	"github.com/ayuxsec/xurl" // go get "github.com/ayuxsec/xurl"
)

func main() {
	rawURLs := []string{
		"https://example.com/path/to/page?user=admin&id=123#section",
		"https://example.com/another/path?debug=true",
	}

	words, errs := xurl.ExtractWords(rawURLs)

	fmt.Print(words) // [ path to page user admin id 123 section another debug true]

	if len(errs) > 0 {
		fmt.Println("\nErrors:")
		for _, err := range errs {
			log.Println(err)
		}
	}
}
```
