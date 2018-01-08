package base

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	CHINAZONE = "Asia/Urumqi"
	UTCZONE   = "UTC"
	KOREAZONE = "Asia/Seoul"
)

func GetChinaTime(timeParam time.Time) (current time.Time) {
	loc, _ := time.LoadLocation(CHINAZONE)
	current = timeParam.In(loc)
	return
}

func GetZoneTime(zone string, timeParam time.Time) (current time.Time) {
	loc, _ := time.LoadLocation(zone)
	current = timeParam.In(loc)
	return
}

func GetDateFormat(timeParam time.Time, format int) (current string) {
	switch format {
	case 112:
		current = timeParam.Format("20060102")
	case 120:
		current = timeParam.Format("20060102 15:04:05")
	case 121:
		current = timeParam.Format("20060102150405")
	case 108:
		current = timeParam.Format("15:04:05")
	default:
		current = timeParam.Format("2006-01-02 15:04:05")
	}
	return
}

func NewDateFormat(t time.Time, format string) (current string) {
	format = strings.ToLower(format)
	switch format {
	case "yymmdd":
		current = fmt.Sprintf("%v%02d%02d", strconv.FormatInt(int64(t.Year()), 10)[2:], t.Month(), t.Day())
	case "yyyymmdd":
		current = t.Format("20060102")
	case "yyyy-mm-dd":
		current = t.Format("2006-01-02")
	case "yyyymmdd hh:mm:ss":
		current = t.Format("20060102 15:04:05")
	case "yyyymmddhhmmss":
		current = t.Format("20060102150405")
	case "hh:mm:ss":
		current = t.Format("15:04:05")
	default:
		current = t.Format("2006-01-02 15:04:05")
	}
	return
}
