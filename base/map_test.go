package base

import (
	"github.com/relax-space/go-kit/test"
	"testing"
)

func Test_JoinMapString(t *testing.T) {

	param := make(map[string]string, 0)
	param["app_id"] = `wx1"234`
	param["name"] = "query"

	newParam := JoinMapString(param)
	test.Equals(t, "app_id=wx1\"234&name=query", newParam)
}

func Test_JoinMapStringEncode(t *testing.T) {

	param := make(map[string]string, 0)
	param["app_id"] = `wx1"234`
	param["name"] = "query"

	newParam := JoinMapStringEncode(param)
	test.Equals(t, "app_id=wx1%22234&name=query", newParam)
}

func Test_JoinMapObject(t *testing.T) {

	param := make(map[string]interface{}, 0)
	param["name"] = "query"
	param["age"] = 18
	param["app_id"] = `wx1"234`

	newParam := JoinMapObject(param)
	test.Equals(t, "age=18&app_id=wx1\"234&name=query", newParam)
}

func Test_JoinMapObjectEncode(t *testing.T) {

	param := make(map[string]interface{}, 0)
	param["name"] = "query"
	param["age"] = 18
	param["app_id"] = `wx1"234`

	newParam := JoinMapObjectEncode(param)
	test.Equals(t, "age=18&app_id=wx1%22234&name=query", newParam)
}
