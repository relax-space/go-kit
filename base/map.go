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

func ParseMapObjectEncode(param string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	if len(param) == 0 {
		return
	}
	vs := strings.Split(param, "&amp;")
	for _, v := range vs {
		if !strings.Contains(v, "=") {
			return
		}
		vss := strings.Split(v, "=")
		vsse, err := url.QueryUnescape(vss[1])
		if err != nil {
			return
		}
		result[vss[0]] = vsse
	}
	return
}

func ParseMapObject(param string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	if len(param) == 0 {
		return
	}
	vs := strings.Split(param, "&")
	for _, v := range vs {
		if !strings.Contains(v, "=") {
			return
		}
		vss := strings.Split(v, "=")
		result[vss[0]] = vss[1]
	}
	return
}
