package jj

import (
	"bytes"
	"fmt"
)

type Attr struct {
	buff *bytes.Buffer
}

func NewAttr() *Attr {
	return &Attr{bytes.NewBuffer([]byte{})}
}

func (a *Attr) Set(k string, v any) *Attr {
	if v == nil {
		a.buff.WriteString(" ")
		a.buff.WriteString(k)
		return a
	}

	a.buff.WriteString(" ")
	a.buff.WriteString(k)
	a.buff.WriteString(`="`)
	a.buff.WriteString(fmt.Sprint(v))
	a.buff.WriteString(`"`)
	return a
}

func (a *Attr) Bytes() []byte {
	return a.buff.Bytes()
}

func (a *Attr) Class(c string) *Attr {
	return a.Set("class", c)
}

func (a *Attr) ID(id string) *Attr {
	return a.Set("id", id)
}

func (a *Attr) Style(s string) *Attr {
	return a.Set("style", s)
}

func (a *Attr) Href(href string) *Attr {
	return a.Set("href", href)
}

func (a *Attr) Target(target string) *Attr {
	return a.Set("target", target)
}

func (a *Attr) Title(title string) *Attr {
	return a.Set("title", title)
}

func (a *Attr) Alt(alt string) *Attr {
	return a.Set("alt", alt)
}

func (a *Attr) Src(src string) *Attr {
	return a.Set("src", src)
}

func (a *Attr) Value(value string) *Attr {
	return a.Set("value", value)
}

func (a *Attr) Placeholder(placeholder string) *Attr {
	return a.Set("placeholder", placeholder)
}

func (a *Attr) Disabled() *Attr {
	return a.Set("disabled", nil)
}

func (a *Attr) Required() *Attr {
	return a.Set("required", nil)
}

func (a *Attr) Readonly() *Attr {
	return a.Set("readonly", nil)
}

func (a *Attr) Checked() *Attr {
	return a.Set("checked", nil)
}

func (a *Attr) Selected() *Attr {
	return a.Set("selected", nil)
}

func (a *Attr) Multiple() *Attr {
	return a.Set("multiple", nil)
}

func (a *Attr) Autofocus() *Attr {
	return a.Set("autofocus", nil)
}

func (a *Attr) Min(min int) *Attr {
	return a.Set("min", min)
}

func (a *Attr) Max(max int) *Attr {
	return a.Set("max", max)
}

func (a *Attr) Step(step int) *Attr {
	return a.Set("step", step)
}

func (a *Attr) Type(t string) *Attr {
	return a.Set("type", t)
}

func (a *Attr) Name(name string) *Attr {
	return a.Set("name", name)
}

func (a *Attr) For(name string) *Attr {
	return a.Set("for", name)
}

func (a *Attr) Rows(rows int) *Attr {
	return a.Set("rows", rows)
}

func (a *Attr) Cols(cols int) *Attr {
	return a.Set("cols", cols)
}

func (a *Attr) Size(size int) *Attr {
	return a.Set("size", size)
}

func (a *Attr) Height(height int) *Attr {
	return a.Set("height", height)
}

func (a *Attr) Width(width int) *Attr {
	return a.Set("width", width)
}

func (a *Attr) Accept(accept string) *Attr {
	return a.Set("accept", accept)
}

func (a *Attr) AcceptCharset(acceptCharset string) *Attr {
	return a.Set("accept-charset", acceptCharset)
}

func (a *Attr) Action(action string) *Attr {
	return a.Set("action", action)
}

func (a *Attr) Autocomplete(autocomplete string) *Attr {
	return a.Set("autocomplete", autocomplete)
}

func (a *Attr) Autoplay(autoplay string) *Attr {
	return a.Set("autoplay", autoplay)
}

func (a *Attr) Capture(capture string) *Attr {
	return a.Set("capture", capture)
}

func (a *Attr) Enctype(enctype string) *Attr {
	return a.Set("enctype", enctype)
}

func (a *Attr) Form(form string) *Attr {
	return a.Set("form", form)
}

func (a *Attr) Method(method string) *Attr {
	return a.Set("method", method)
}

func (a *Attr) NoValidate() *Attr {
	return a.Set("novalidate", nil)
}

func (a *Attr) Pattern(pattern string) *Attr {
	return a.Set("pattern", pattern)
}
