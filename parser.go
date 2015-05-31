package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Mapped map[string]interface{}

type Parser struct {
	Resp        *http.Response
	ContentType string
}

func (p *Parser) GetContent() {
	contentType := p.Resp.Header.Get("Content-Type")
	p.ContentType = strings.Split(contentType, ";")[0]
}

func (p Parser) ParseBody() []Mapped {
	var mapping []Mapped
	defer p.Resp.Body.Close()

	if p.ContentType == "application/json" {
		b, err := ioutil.ReadAll(p.Resp.Body)

		if err != nil {
			log.Println(err)

			os.Exit(1)
		}
		json.Unmarshal(b, &mapping)
	}
	// CSV not implemented

	return mapping
}
