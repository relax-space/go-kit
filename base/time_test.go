package base

import (
	"fmt"
	"go-kit/test"
	"testing"
	"time"
)

func Test_GetChinaTime(t *testing.T) {
	now := time.Now()
	newNow := GetChinaTime(now)
	test.Equals(t, newNow.Format("2006-01-02 15:04:05"), now.Format("2006-01-02 15:04:05"))
}

func Test_GetZoneTime(t *testing.T) {
	now := time.Now()
	utcTime := GetZoneTime(UTCZONE, now)
	chinaTime := GetZoneTime(CHINAZONE, now)
	test.Equals(t, utcTime.Add(8*time.Hour).Format("2006-01-02 15:04:05"), chinaTime.Format("2006-01-02 15:04:05"))
}

func Test_GetDateFormat(t *testing.T) {
	now := time.Now()
	d112 := GetDateFormat(now, 112)
	d120 := GetDateFormat(now, 120)
	d121 := GetDateFormat(now, 121)
	d108 := GetDateFormat(now, 108)
	d0 := GetDateFormat(now, 0)

	fmt.Println(d112)
	fmt.Println(d120)
	fmt.Println(d121)
	fmt.Println(d108)
	fmt.Println(d0)

	test.Equals(t, len(d112), 8)
	test.Equals(t, len(d120), 17)
	test.Equals(t, len(d121), 14)
	test.Equals(t, len(d108), 8)
	test.Equals(t, len(d0), 19)

}
