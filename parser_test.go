package parser

import (
	"net/http"
	"testing"
)

var (
	jsonContent = "http://www.mocky.io/v2/556b0ba673eedce302329da9"
	jsonExtraContent = "http://www.mocky.io/v2/556b214973eedc9503329dae"
	csvContent  = "http://www.mocky.io/v2/556b0bf273eedceb02329daa"
	csvExtraContent = "http://www.mocky.io/v2/556b1fbc73eedc7e03329dad"
)

func makeRequest(reqType string, extra bool) string {
	if reqType == "csv" {
		if extra {
			return csvExtraContent
		}
		return csvContent
	}

	if extra {
		return jsonExtraContent
	}
	return jsonContent
}

func TestGetJsonContent(t *testing.T) {
	resps, _ := http.Get(makeRequest("json", false))

	expected := "application/json"
	parser := &Parser{Resp: resps}
	parser.GetContent()
	got := parser.ContentType

	if expected != got {
		t.Errorf("Should be %s, returned %s", expected, got)
	}
}

func TestGetCsvContent(t *testing.T) {
	resps, _ := http.Get(makeRequest("csv", false))

	expected := "text/csv"
	parser := &Parser{Resp: resps}
	parser.GetContent()
	got := parser.ContentType

	if expected != got {
		t.Errorf("Should be %s, returned %s", expected, got)
	}
}

func TestParseJsonBody(t *testing.T) {
	resps, _ := http.Get(makeRequest("json", false))
	expectedFields := []string{"name", "age", "email", "sex"}

	parser := &Parser{Resp: resps}
	parser.GetContent()
	mapped := parser.ParseBody()

	for _, field := range expectedFields {
		if _, ok := mapped[0][field]; !ok {
			t.Errorf("Key %s not found", field)
		}
	}
}

func TestParseJsonExtraBody(t *testing.T) {
	resps, _ := http.Get(makeRequest("json", true))
	expectedFields := []string{"name", "age", "email", "sex", "number", "other"}

	parser := &Parser{Resp: resps}
	parser.GetContent()
	mapped := parser.ParseBody()

	for _, field := range expectedFields {
		if _, ok := mapped[0][field]; !ok {
			t.Errorf("Key %s not found", field)
		}
	}
}
