package sign

import (
	"kit/test"
	"testing"
)

func Test_GetMD5Hash(t *testing.T) {
	_, err := GetMD5Hash("secret jwt")
	test.Ok(t, err)
	_, err = GetMD5Hash("secret jwt", true)
	test.Ok(t, err)

}

func Test_CheckSign(t *testing.T) {
	url := "this is a secret"
	key := "20160101"
	sign := MakeMd5Sign(url, key)
	test.Assert(t, CheckSign(url, key, sign), "check sign failure")
}
