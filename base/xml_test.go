package base

import (
	"fmt"
	"go-kit/test"
	"testing"
)

func Test_XmlToMap(t *testing.T) {
	xmlData := `<xml>
	<app_id>wx1111</app_id>
	<name>apple</name>
	</xml>`

	data, err := XmlToMap(xmlData)
	test.Ok(t, err)
	fmt.Printf("\n%+v", data)
	fmt.Printf("\n%+v", data["xml"])

}

func Test_XmlToMap_2(t *testing.T) {
	xmlData := `<apple>
	<color>red</color>
	<price>18</price>
	</apple>`

	data, err := XmlToMap(xmlData)
	test.Ok(t, err)
	fmt.Printf("\n%+v", data)
	fmt.Printf("\n%+v", data["apple"])

}
