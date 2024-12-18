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

func Loop(j *jj.J, users []string) {
	j.Ul(nil, func() {
		for _, user := range users {
			j.Li(nil, func() {
				j.Text(user)
			})
		}
	})
}
