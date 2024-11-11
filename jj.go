package jj

import (
	"fmt"
	"html"
	"io"
)

type J struct {
	io.Writer
}

type JFunc func(J)
type Attrs map[string]any

func New(w io.Writer) J {
	return J{w}
}

func (j J) write(str string) {
	j.Write([]byte(str))
}

func (j J) element(tag string, attrs Attrs, f JFunc) {
	j.write("<")
	j.write(tag)
	j.write(j.attr(attrs))
	j.write(">")
	f(j)
	j.write("</")
	j.write(tag)
	j.write(">")
}

func (j J) voidElement(text string, attrs Attrs) {
	j.write("<")
	j.write(text)
	j.write(j.attr(attrs))
	j.write("/>")
}

func (j J) attr(attrs Attrs) string {
	result := ""
	for k, v := range attrs {
		if v == nil {
			result += fmt.Sprintf(" %s", k)
		} else {
			result += fmt.Sprintf(" %s=\"%v\"", k, v)
		}
	}

	return result
}

func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func (j J) CustomElement(tag string, attrs Attrs, f JFunc) {
	j.element(tag, attrs, f)
}

func (j J) CustomVoidElement(tag string, attrs Attrs) {
	j.voidElement(tag, attrs)
}

func (j J) Text(text string) {
	j.write(html.EscapeString(text))
}

// Raw writes raw HTML to the output. be careful with this.
func (j J) Raw(html string) {
	j.write(html)
}
