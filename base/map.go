package base

import (
	"bytes"
	"encoding/json"
	"net/url"
	"sort"
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
