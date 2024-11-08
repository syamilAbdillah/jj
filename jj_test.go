package jj

import (
	"bytes"
	"fmt"
	"testing"
)

func TestJ(t *testing.T) {
	var b bytes.Buffer
	New(&b).CustomElement("div", nil, func(j J) {
		j.CustomElement("span", nil, func(j J) {
			j.Text("Hello")
		})
	})
	if b.String() != "<div><span>Hello</span></div>" {
		t.Errorf("Expected <div><span>Hello</span></div>, got %s", b.String())
	}
}

func TestJVoidElement(t *testing.T) {
	var b bytes.Buffer
	New(&b).CustomVoidElement("div", nil)
	if b.String() != "<div/>" {
		t.Errorf("Expected <div/>, got %s", b.String())
	}
}

func TestChildrenNodes(t *testing.T) {
	var b bytes.Buffer
	New(&b).CustomElement("div", nil, func(j J) {
		j.CustomElement("span", nil, func(j J) {
			j.Text("Hello")
		})
		j.CustomElement("span", nil, func(j J) {
			j.Text("World")
		})
	})
	expected := "<div><span>Hello</span><span>World</span></div>"
	if b.String() != expected {
		t.Errorf("Expected %s, got %s", expected, b.String())
	}
}

func TestJText(t *testing.T) {
	var b bytes.Buffer
	New(&b).Text("Hello")
	if b.String() != "Hello" {
		t.Errorf("Expected Hello, got %s", b.String())
	}
}

func TestMultipleAttributes(t *testing.T) {
	var b bytes.Buffer
	New(&b).CustomElement("div", Attrs{"class": "foo", "id": "bar"}, func(j J) {
		j.Text("Hello")
	})
	if b.String() != "<div class=\"foo\" id=\"bar\">Hello</div>" || b.String() != "<div id=\"bar\" class=\"foo\">Hello</div>" {
		t.Errorf("Expected <div class=\"foo\" id=\"bar\">Hello</div>, got %s", b.String())
	}
}

func TestLooping(t *testing.T) {
	var b bytes.Buffer
	New(&b).Div(nil, func(j J) {
		for i := 0; i < 10; i++ {
			j.Span(nil, func(j J) {
				j.Text(fmt.Sprintf("%d", i))
			})
		}
	})
	expected := "<div><span>0</span><span>1</span><span>2</span><span>3</span><span>4</span><span>5</span><span>6</span><span>7</span><span>8</span><span>9</span></div>"
	if b.String() != expected {
		t.Errorf("Expected %s, got %s", expected, b.String())
	}
}

func TestNilAttributes(t *testing.T) {
	var b bytes.Buffer
	New(&b).Input(Attrs{"required": nil})
	if b.String() != "<input required/>" {
		t.Errorf("Expected <input required/> got %s", b.String())
	}
}
