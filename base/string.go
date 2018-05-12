package base

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func ToString(raw interface{}) string {
	switch v := raw.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(uint64(v), 10)
	case bool:
		return strconv.FormatBool(bool(v))

	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case string:
		return string(v)
	}
	val := reflect.ValueOf(raw)
	switch val.Kind() {
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice, reflect.Ptr:
		b, _ := json.Marshal(raw)
		return string(b)
	}
	return ""
}

//b contains a
func ContainsA(a []string, b string) bool {
	for _, v := range a {
		if strings.Contains(b, v) {
			return true
		}
	}
	return false
}

//a contains b
func ContainsB(a []string, b string) bool {
	for _, v := range a {
		if strings.Contains(v, b) {
			return true
		}
	}
	return false
}

func Contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

/*
rawString:	name=xiao.xinmiao&age=18&sign=3F7EC1885326B9D1FD078DB2276C84D6
key:		sign
after:		name=xiao.xinmiao&age=18
*/
func RemoveFromString(rawString, key string) (result string) {
	if !strings.Contains(rawString, key) {
		return
	}
	startIndex := strings.Index(rawString, key+"=")
	tempParam := rawString[startIndex:]
	secordIndex := strings.Index(tempParam, "&")
	var secondPart string
	if secordIndex != -1 {
		secondPart = tempParam[secordIndex:]
	}
	result = rawString[:startIndex-1] + secondPart
	return
}
