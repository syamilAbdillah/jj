package jj

import (
	"fmt"
	"html"
	"io"
)

type J struct {
	w io.Writer
}

func (j *J) writeString(str string) {
	j.w.Write([]byte(str))
}

type Attributes interface {
	Bytes() []byte
}

func (j *J) open(tag string, attr Attributes) {
	j.writeString("<")
	j.writeString(tag)
	if attr != nil {
		j.w.Write(attr.Bytes())
	}
	j.writeString(">")
}

func (j *J) close(tag string) {
	j.writeString("</")
	j.writeString(tag)
	j.writeString(">")
}

func (j *J) Element(tag string, attrs Attributes, f func()) {
	j.open(tag, attrs)

	if f != nil {
		f()
		j.close(tag)
	}
}

func Render(w io.Writer, f func(j *J)) {
	j := J{w}
	f(&j)
}

func (j *J) Text(text string) {
	j.writeString(html.EscapeString(text))
}

func (j *J) Textf(f string, a ...any) {
	t := fmt.Sprintf(f, a...)
	j.writeString(t)
}

func (j *J) TextFunc(t string) func() {
	return func() {
		j.Text(t)
	}
}

func (j *J) TextfFunc(f string, a ...any) func() {
	return func() {
		j.Textf(f, a...)
	}
}

// Raw writes raw HTML to the output. be careful with this.
func (j *J) Raw(html string) {
	j.writeString(html)
}
