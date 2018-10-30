package render

import (
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	templatePath = "../testdata"
)

func TestSetTemplatePath(t *testing.T) {
	SetTemplatePath(templatePath)

	got := GetTemplatePath()
	want := templatePath
	if got != want {
		t.Errorf("Got %v, Want %v", got, want)
	}
}

func TestHTML(t *testing.T) {
	w := httptest.NewRecorder()

	SetTemplatePath(templatePath)

	HTML(w, "test.html", struct{}{})

	got := w.Body.String()
	want := "<h1>test</h1>"

	if got != want {
		t.Errorf("Got %v, Want %v", got, want)
	}
}

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()

	data := struct {
		Name string `json:"name"`
	}{
		"choto",
	}

	JSON(w, data)

	cases := []struct {
		Got  string
		Want string
	}{
		{strings.TrimSpace(w.Body.String()), `{"name":"choto"}`},
		{w.Header().Get("Content-Type"), "application/json; charset=utf-8"},
	}

	for _, c := range cases {
		if c.Got != c.Want {
			t.Errorf("Got %v, Want %v\n", c.Got, c.Want)
		}
	}
}
