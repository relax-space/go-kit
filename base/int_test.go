package base

import (
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_ContainsInt(t *testing.T) {
	var act int64 = 45
	test.Assert(t, ContainsInt([]int64{1, 2, 3, 45, 6}, act), "data("+ToString(act)+") cannot be found")
}
func Test_ContainsInts(t *testing.T) {
	var act int64 = 45
	test.Assert(t, ContainsInts("1, 2, 3, 45, 6", act), "data("+ToString(act)+") cannot be found")
}
