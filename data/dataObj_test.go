package data

import (
	"kit/test"
	"testing"
)

func Test_FromXml(t *testing.T) {

	xmlData := `<xml>
	<app_id>wx1111</app_id>
	<name>apple</name>
	</xml>`
	atomData := NewDataObj()
	test.Ok(t, atomData.FromXml(xmlData))

	test.Equals(t, atomData.DataAttr["app_id"], "wx1111")

}

func Test_ToXml_ToUrl_ToJson(t *testing.T) {
	xmlData := `<xml>
 <app_id>wx1111</app_id>
 <name>apple</name>
 </xml>`
	atomData := NewDataObj()
	test.Ok(t, atomData.FromXml(xmlData))

	test.Equals(t, `{"app_id":"wx1111","name":"apple"}`, atomData.ToJson())
	test.Equals(t, "app_id=wx1111&name=apple", atomData.ToUrl())

	xml01 := `<xml>
 <app_id>wx1111</app_id>
 <name>apple</name>
</xml>`
	xml02 := `<xml>
 <name>apple</name>
 <app_id>wx1111</app_id>
</xml>`
	got := atomData.ToXml()
	test.Assert(t, got == xml01 || got == xml02, "xml transfer failure")
}
