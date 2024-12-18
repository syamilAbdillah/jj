package jj

import (
	"bytes"
	"testing"
)

func TestJ(t *testing.T) {
	var b bytes.Buffer
	Render(&b, func(j *J) {
		j.Element("div", nil, func() {
			j.Element("span", nil, func() {
				j.Text("Hello")
			})
		})
	})
	if b.String() != "<div><span>Hello</span></div>" {
		t.Errorf("Expected <div><span>Hello</span></div>, got %s", b.String())
	}
}

func TestJVoidElement(t *testing.T) {
	var b bytes.Buffer
	Render(&b, func(j *J) {
		j.Element("br", nil, nil)
	})
	if b.String() != "<br>" {
		t.Errorf("Expected <br>, got %s", b.String())
	}
}

func TestChildrenNodes(t *testing.T) {
	var b bytes.Buffer
	Render(&b, func(j *J) {
		j.Div(nil, func() {
			j.Span(nil, func() {
				j.Text("Hello")
			})
			j.Span(nil, func() {
				j.Text("World")
			})
		})
	})
	expected := "<div><span>Hello</span><span>World</span></div>"
	if b.String() != expected {
		t.Errorf("Expected %s, got %s", expected, b.String())
	}
}

func TestJText(t *testing.T) {
	var b bytes.Buffer
	Render(&b, func(j *J) {
		j.Text("Hello")
	})
	if b.String() != "Hello" {
		t.Errorf("Expected Hello, got %s", b.String())
	}
}

func TestMultipleAttributes(t *testing.T) {
	var b bytes.Buffer
	Render(&b, func(j *J) {
		j.Div(NewAttr().Class("foo").ID("bar"), func() {
			j.Text("Hello")
		})
	})
	if b.String() != "<div class=\"foo\" id=\"bar\">Hello</div>" {
		t.Errorf("Expected <div class=\"foo\" id=\"bar\">Hello</div>, got %s", b.String())
	}
}

func TestLooping(t *testing.T) {
	var b bytes.Buffer
	Render(&b, func(j *J) {
		j.Div(nil, func() {
			for i := 0; i < 10; i++ {
				j.Span(nil, j.TextfFunc("%d", i))
			}
		})
	})
	expected := "<div><span>0</span><span>1</span><span>2</span><span>3</span><span>4</span><span>5</span><span>6</span><span>7</span><span>8</span><span>9</span></div>"
	if b.String() != expected {
		t.Errorf("Expected %s, got %s", expected, b.String())
	}
}

func TestVoidAttributes(t *testing.T) {
	var b bytes.Buffer
	Render(&b, func(j *J) {
		j.Input(NewAttr().Required())
	})

	expected := "<input required>"
	if b.String() != expected {
		t.Errorf("Expected %s got %s", expected, b.String())
	}
}
