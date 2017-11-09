package mapstruct

import (
	"fmt"
	"kit/test"
	"testing"
)

func Test_Decode(t *testing.T) {

	var apple struct {
		Color string `json:"color"`
		Price int    `json:"price"`
	}

	appleMap := map[string]interface{}{
		"color": "red",
		"price": 14,
	}

	test.Ok(t, Decode(appleMap, &apple))
	fmt.Printf("%+v", apple)

	test.Equals(t, "red", apple.Color)

}
