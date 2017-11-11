package data

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/relax-space/go-kit/base"
)

type Data struct {
	DataAttr map[string]interface{}
}

func New() *Data {
	return &Data{
		DataAttr: map[string]interface{}{},
	}
}

func (a *Data) ToXml() string {
	if len(a.DataAttr) == 0 {
		return ""
	}
	xmlMap := base.XmlMap{}
	for k, v := range a.DataAttr {
		xmlMap[k] = v
	}
	if x, err := xml.MarshalIndent(xmlMap, "", " "); err == nil {
		return string(x)
	} else {
		fmt.Println(err)

	}
	return ""
}

func (a *Data) ToUrl() string {
	return base.JoinMapObjectEncode(a.DataAttr)
}

func (a *Data) ToJson() string {
	v, _ := json.Marshal(a.DataAttr)
	return string(v)
}

func (a *Data) FromXml(xmlStr string) (err error) {
	var data map[string]interface{}
	data, err = base.XmlToMap(xmlStr)
	if err != nil {
		return
	}
	v, has := data["xml"]
	if has {
		a.DataAttr = v.(map[string]interface{})
	}
	return
}

func (a *Data) SetValue(key string, value interface{}) {
	a.DataAttr[key] = value
}

func (a *Data) GetValue(key string) interface{} {
	return a.DataAttr[key]
}
func (a *Data) RemoveKey(key string) {
	delete(a.DataAttr, key)
}

func (a *Data) IsSet(key string) bool {
	_, ok := a.DataAttr[key]
	return ok
}
