package base

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_MarshalXML(t *testing.T) {
	xmlMap := XmlMap{}
	xmlMap["color"] = "red"
	xmlMap["price"] = 18

	s, err := xml.MarshalIndent(xmlMap, "", " ")
	fmt.Println(string(s))

	test.Ok(t, err)

}
