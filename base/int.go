package base

import (
	"strconv"
	"strings"
)

func ContainsInt(list []int64, value int64) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsInts(list string, value int64) (has bool) {
	slist := strings.Split(list, ",")
	for _, v := range slist {
		var v1 int64
		v1, _ = strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if v1 == value {
			has = true
			return
		}
	}
	return
}
