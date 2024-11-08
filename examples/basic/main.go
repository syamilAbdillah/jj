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
