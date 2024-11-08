# JJ - A Go HTML Builder

Try HTML components in pure Go.

## Install

```sh
go get github.com/syamilAbdillah/jj
```

> don't install this shit. just copy and paste it to your codebase.
>
> shamelessly stole [htmgolang](https://github.com/htmgolang/htmg) code and changed it to my heart's content.

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

#### Conditional

```go

func Conditional(j jj.J, isLoggedIn bool) {
	href := "/login"
	if isLoggedIn {
		href = "/profile"
	}
	j.A(jj.Attrs{"href": href}, func(j jj.J) {
		if isLoggedIn {
			j.Text("Profile")
		} else {
			j.Text("Login")
		}
	})
}
```

#### Loop

```go

func Loop(j jj.J, users []string) {
	j.Ul(nil, func(j jj.J) {
		for _, user := range users {
			j.Li(nil, func(j jj.J) {
				j.Text(user)
			})
		}
	})
}
```

## TODO

- <input type="checkbox" disabled /> codegen tool from HTML to JJ syntax
- <input type="checkbox" disabled /> more tests
- <input type="checkbox" disabled /> more examples
- <input type="checkbox" disabled /> documentation site

### why the name is JJ?

because my index finger is on `j` key

## License

MIT
