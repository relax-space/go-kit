package base

import (
	"bytes"
	"encoding/json"
	"net/url"
	"sort"
	"strings"
)

func JoinMapString(v map[string]string) string {

	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if len(vs) == 0 {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(vs)
	}
	return buf.String()
}

func JoinMapStringEncode(v map[string]string) string {

	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if len(vs) == 0 {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(vs))

	}
	return buf.String()
}
func JoinMapObject(v map[string]interface{}) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if vs == nil {
			continue
		}
		if str, ok := vs.(string); ok {
			if len(str) == 0 {
				continue
			}
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		if tValue, ok := vs.(string); ok {
			buf.WriteString(tValue)

		} else if tValue, ok := vs.([]byte); ok {
			buf.Write(tValue)
		} else {
			b, _ := json.Marshal(vs)
			buf.Write(b)
		}
	}
	return buf.String()
}

func JoinMapObjectEncode(v map[string]interface{}) string {

	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if vs == nil {
			continue
		}
		if str, ok := vs.(string); ok {
			if len(str) == 0 {
				continue
			}
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		if tValue, ok := vs.(string); ok {
			buf.WriteString(url.QueryEscape(tValue))
		} else if tValue, ok := vs.([]byte); ok {
			buf.Write(tValue)
		} else {
			b, _ := json.Marshal(vs)
			buf.Write(b)
		}
	}
	return buf.String()
}

//sep1: default value & ,sep2:default value =
func ParseMapObjectEncode(param, sep1, sep2 string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	if len(param) == 0 {
		return
	}
	vs := strings.Split(param, sep1)
	for _, v := range vs {
		if !strings.Contains(v, sep2) {
			return
		}
		vss := strings.Split(v, sep2)
		vsse, err := url.QueryUnescape(vss[1])
		if err != nil {
			return
		}
		result[vss[0]] = vsse
	}
	return
}

//sep1: default value & ,sep2:default value =
func ParseMapObject(param, sep1, sep2 string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	if len(param) == 0 {
		return
	}
	vs := strings.Split(param, sep1)
	for _, v := range vs {
		if !strings.Contains(v, sep2) {
			return
		}
		vss := strings.Split(v, sep2)
		result[vss[0]] = vss[1]
	}
	return
}

func JoinMap(v map[string]interface{}, sep1, sep2 string) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if vs == nil {
			continue
		}
		if str, ok := vs.(string); ok {
			if len(str) == 0 {
				continue
			}
		}
		if buf.Len() > 0 {
			buf.WriteString(sep1) //'&'
		}
		buf.WriteString(k)
		buf.WriteString(sep2) //'='
		if tValue, ok := vs.(string); ok {
			buf.WriteString(tValue)

		} else if tValue, ok := vs.([]byte); ok {
			buf.Write(tValue)
		} else {
			b, _ := json.Marshal(vs)
			buf.Write(b)
		}
	}
	return buf.String()
}

func JoinMapEncode(v map[string]interface{}, sep1, sep2 string) string {

	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if vs == nil {
			continue
		}
		if str, ok := vs.(string); ok {
			if len(str) == 0 {
				continue
			}
		}
		if buf.Len() > 0 {
			buf.WriteString(sep1) //'&'
		}

		buf.WriteString(url.QueryEscape(k))
		buf.WriteString(sep2) //'='
		if tValue, ok := vs.(string); ok {
			buf.WriteString(url.QueryEscape(tValue))
		} else if tValue, ok := vs.([]byte); ok {
			buf.Write(tValue)
		} else {
			b, _ := json.Marshal(vs)
			buf.Write(b)
		}
	}
	return buf.String()
}

func JoinMapJsonRawMessage(v map[string]*json.RawMessage) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if vs == nil {
			continue
		}
		vsTrim := strings.TrimPrefix(string(*vs), "\"")
		vsTrim = strings.TrimSuffix(vsTrim, "\"")
		values := string(vsTrim)
		if len(values) == 0 {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(values)
	}
	return buf.String()
}

//to be: + <=> %20
func JoinMapJsonRawMessageEncode(v map[string]*json.RawMessage) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if vs == nil {
			continue
		}
		vsTrim := strings.TrimPrefix(string(*vs), "\"")
		vsTrim = strings.TrimSuffix(vsTrim, "\"")
		values := string(vsTrim)
		if len(values) == 0 {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(values))
	}
	return buf.String()
}
