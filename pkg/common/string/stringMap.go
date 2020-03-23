package string

import (
	"encoding/xml"
	"fmt"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorStringAVA "github.com/ver13/ava/pkg/common/string/error"
)

// StringMap is a map[string]string.
type StringMap map[string]string

// StringMap marshals into XML.
func (s StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) *errorAVA.Error {

	tokens := []xml.Token{start}

	for key, value := range s {
		t := xml.StartElement{Name: xml.Name{Local: key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{Name: t.Name})
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return errorStringAVA.EncodeTokenXml(err, fmt.Sprintf("%v", t))
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return errorStringAVA.FlushTokenXml(err, fmt.Sprintf("%v", e))
	}

	return nil
}

func (s StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) *errorAVA.Error {
	key := ""
	val := ""

	for {
		t, _ := d.Token()
		switch tt := t.(type) {

		// TODO: parse the inner structure

		case xml.StartElement:
			fmt.Println(">", tt)

		case xml.EndElement:
			fmt.Println("<", tt)
			if tt.Name == start.Name {
				return nil
			}

			if tt.Name.Local == "enabled" {
				s[key] = val
			}
		}
	}
}
