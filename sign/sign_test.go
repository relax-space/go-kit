package sign

import (
	"fmt"
	"kit/test"
	"testing"
)

func Test_GetMD5Hash(t *testing.T) {
	s, err := GetMD5Hash("secret jwt")

	test.Ok(t, err)
	s1, err := GetMD5Hash("secret jwt", true)
	test.Ok(t, err)
	fmt.Println(s, s1)

}

func Test_CheckSign(t *testing.T) {
	url := "this is a secret"
	key := "20160101"
	sign := MakeMd5Sign(url, key)
	test.Assert(t, CheckMd5Sign(url, key, sign), "check sign failure")
}

func Test_MakeSha1Sign(t *testing.T) {
	priKey := `MIICXQIBAAKBgQDuRElXZylWgZj8awrucgjHT0wzNoVfZ4qUfyQFPAM0okbWlCY1B5jrl3W+AYjQnWT139r8i7RVcWnVDX8ZkELhqX81/yPdA4Yu+Eay4BaPP0/lkcuAZ3FWKAdiFIQRUVUElTz4qeDa2WHKzhDqCoKFsNd9tZ2H6FsQbrbt7GNXJwIBAwKBgQCe2DDk73DkVmX9nLH0TAXaNN13ea4/mlxi/21Y0qzNwYSPDW7OBRCdD6PUAQXgaO35P+dTB82OS5vjXlS7tYHqfPitAm65syhhC+aSfjaMd7n5oawKMp58Z0XEd5VJVuAXTgOvkTOi8Tc1QXK1IMnVI9Dg1C5ahfj9Y5igapjmmwJBAPvE1P7/jcWotnyWXZJ/51exADBZQdrOMoRh0YAfDckcet2+e1pQU0ohuAcJHm1qGp1nhF5eAhukymZY2wFZnbsCQQDyRV18fjixIObp1nmQRNU0B+7u8C9AtWkIvOnhBRWF6GbCePX1wxMlbULk27vr6atdtqgY+f4DwMny+SJLJF+FAkEAp9iN/1UJLnB5qGQ+YaqaOnYAIDuBPIl3AuvhABSz22hR6SmnkYriMWvQBLC+85wRvkUC6ZQBZ8MxmZCSAOZpJwJBAKGDk6hUJctrRJvkUQrYjiKv9J9Kyisjm1som+tYuQPwRIGl+U6CDMOeLJiSfUfxHOkkcBCmqVfV2/dQwYdtlQMCQQD3sPdGVBf0EQgi0TaalCUi9XiZ9DL650Iw1EP7jBTmNd2Fc6u2iXwPpFJ6tu9D9nb4Kxw6IZi+4XitLZdJBrQ2`
	text := "123456"
	signStr, err := MakeSha1Sign(text, priKey)
	fmt.Println(signStr)
	test.Ok(t, err)
}

//Don't worry about security, this public and private key is testing key
func Test_CheckSha1Sign(t *testing.T) {
	priKey := `MIICXQIBAAKBgQDuRElXZylWgZj8awrucgjHT0wzNoVfZ4qUfyQFPAM0okbWlCY1B5jrl3W+AYjQnWT139r8i7RVcWnVDX8ZkELhqX81/yPdA4Yu+Eay4BaPP0/lkcuAZ3FWKAdiFIQRUVUElTz4qeDa2WHKzhDqCoKFsNd9tZ2H6FsQbrbt7GNXJwIBAwKBgQCe2DDk73DkVmX9nLH0TAXaNN13ea4/mlxi/21Y0qzNwYSPDW7OBRCdD6PUAQXgaO35P+dTB82OS5vjXlS7tYHqfPitAm65syhhC+aSfjaMd7n5oawKMp58Z0XEd5VJVuAXTgOvkTOi8Tc1QXK1IMnVI9Dg1C5ahfj9Y5igapjmmwJBAPvE1P7/jcWotnyWXZJ/51exADBZQdrOMoRh0YAfDckcet2+e1pQU0ohuAcJHm1qGp1nhF5eAhukymZY2wFZnbsCQQDyRV18fjixIObp1nmQRNU0B+7u8C9AtWkIvOnhBRWF6GbCePX1wxMlbULk27vr6atdtqgY+f4DwMny+SJLJF+FAkEAp9iN/1UJLnB5qGQ+YaqaOnYAIDuBPIl3AuvhABSz22hR6SmnkYriMWvQBLC+85wRvkUC6ZQBZ8MxmZCSAOZpJwJBAKGDk6hUJctrRJvkUQrYjiKv9J9Kyisjm1som+tYuQPwRIGl+U6CDMOeLJiSfUfxHOkkcBCmqVfV2/dQwYdtlQMCQQD3sPdGVBf0EQgi0TaalCUi9XiZ9DL650Iw1EP7jBTmNd2Fc6u2iXwPpFJ6tu9D9nb4Kxw6IZi+4XitLZdJBrQ2`
	text := "123456"
	signStr, err := MakeSha1Sign(text, priKey)
	test.Ok(t, err)

	pubKey := `MIGdMA0GCSqGSIb3DQEBAQUAA4GLADCBhwKBgQDuRElXZylWgZj8awrucgjHT0wzNoVfZ4qUfyQFPAM0okbWlCY1B5jrl3W+AYjQnWT139r8i7RVcWnVDX8ZkELhqX81/yPdA4Yu+Eay4BaPP0/lkcuAZ3FWKAdiFIQRUVUElTz4qeDa2WHKzhDqCoKFsNd9tZ2H6FsQbrbt7GNXJwIBAw==`
	test.Assert(t, CheckSha1Sign(text, signStr, pubKey), "check sign failure")

}
