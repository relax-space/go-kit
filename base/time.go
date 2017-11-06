package base

import "time"

const (
	CHINAZONE = "Asia/Shanghai"
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
