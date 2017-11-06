package base

import (
	"encoding/xml"
)

type XmlMap map[string]interface{}

// Map marshals into XML.
func (x XmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "xml"
	tokens := []xml.Token{start}
	for key, value := range x {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(ToString(value)), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}
