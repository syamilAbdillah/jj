# JJ - A Go HTML Builder

Write HTML without leaving your lovely Go files.

## Install

```sh
go get github.com/syamilAbdillah/jj
```

> v2 is break everything, so good luck with that
>
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
		jj.Render(w, func(j *jj.J) {
			j.Raw("<!DOCTYPE html>")
			j.Html(nil, func() {
				j.Head(nil, func() {
					j.Title(nil, func() {
						j.Text("Hello World")
					})
				})
				j.Body(jj.NewAttr().Class("container"), func() {
					j.H1(nil, func() {
						j.Text("Hello World")
					})

					j.Br(nil)
				})
			})
		})
	})

	http.ListenAndServe(":8080", nil)
}

```

#### Conditional

```go


func Conditional(j *jj.J, isLoggedIn bool) {
	atr := jj.NewAttr()
	if isLoggedIn {
		atr.Href("/profile")
	} else {
		atr.Href("/login")
	}

	j.A(atr, func() {
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


func Loop(j *jj.J, users []string) {
	j.Ul(nil, func() {
		for _, user := range users {
			j.Li(nil, func() {
				j.Text(user)
			})
		}
	})
}
```

## TODO

- [ ] codegen tool from HTML to JJ syntax
- [x] more tests
- [x] more examples
- [ ] documentation site

### why the name is JJ?

because my index finger is on `j` key

## License

MIT
