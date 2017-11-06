package auth

import (
	"fmt"
	"kit/test"
	"testing"
)

func Test_NewToken(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	fmt.Println(token)
	test.Ok(t, err)
}

func Test_Extract(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	test.Ok(t, err)
	mapClaims, err := Extract(token)
	test.Ok(t, err)
	test.Equals(t, "wx1234", mapClaims["app_id"])
}

func Test_Renew(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	test.Ok(t, err)
	token, err = Renew(token)
	test.Ok(t, err)
}

func Test_EditPayload(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	test.Ok(t, err)
	addParam := map[string]interface{}{"name": "limon"}
	_, err = EditPayload(token, addParam)
	test.Ok(t, err)
}

func Test_EditPayloadAdd(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	test.Ok(t, err)
	addParam := map[string]interface{}{"name": "limon"}
	token, err = EditPayload(token, addParam)
	test.Ok(t, err)
	claim, err := Extract(token)
	test.Assert(t, "limon" == claim["name"], "add token failure")
	test.Assert(t, "wx1234" == claim["app_id"], "add token failure")
}

func Test_EditPayloadUpdate(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	test.Ok(t, err)
	addParam := map[string]interface{}{"app_id": "wx1111"}
	token, err = EditPayload(token, addParam)
	test.Ok(t, err)
	claim, err := Extract(token)
	test.Assert(t, "wx1111" == claim["app_id"], "edit token failure")
}
