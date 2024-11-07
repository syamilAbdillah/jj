# JJ - A Go HTML Builder

Try HTML components in pure Go.

## Install

```sh
go get github.com/syamilAbdillah/jj
```

don't install this shit. just copy and paste it to your codebase.

shamelessly stole [htmgolang](github.com/htmgolang/htmg) code and changed it to my heart's content.

## features

- Pure Go
- embrace [LOB principle](https://htmx.org/essays/locality-of-behaviour)
- No dependencies
- No reflection
- No Money
- No Girl friend (WTF ... ?)

## Usage

```go
package main

import (
	"net/http"
	"github.com/syamilAbdillah/jj"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		j := jj.New(w)
    j.Raw("<!DOCTYPE html>")
    j.Html(nil, func(j jj.J) {
      j.Head(nil, func(j jj.J) {
        j.Title(nil, func(j jj.J) {
          j.Text("Hello World")
        })
      })
      j.Body(jj.Attrs{"class": "container"}, func(j jj.J) {
        j.H1(nil, func(j jj.J) {
          j.Text("Hello World")
        })
      })
    })
	})

	http.ListenAndServe(":8080", nil)
}
```

## TODO

- [ ] codegen tool from HTML to JJ syntax
- [ ] tests
- [ ] documentation site
- [ ] examples

### why the name is JJ?

because my index finger is on `j` key

## License

MIT
