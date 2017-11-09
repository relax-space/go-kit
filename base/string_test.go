package base

import (
	"fmt"
	"testing"
)

func Test_ToString(t *testing.T) {
	var a int64 = 12
	var b int32 = 13
	var c bool = false
	var d float32 = 14
	var e float64 = 15

	fmt.Printf("\nint64:%T-%v", ToString(a), a)
	fmt.Printf("\nint32:%T-%v", ToString(b), b)
	fmt.Printf("\nbool:%T-%v", ToString(c), c)
	fmt.Printf("\nfloat32:%T-%v", ToString(d), d)
	fmt.Printf("\nfloat64:%T-%v", ToString(e), e)

}
