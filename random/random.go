package random

import (
	"bytes"
	"encoding/binary"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func UuIdInt64() (uuId uint64) {
	u1 := uuid.NewV4()
	buf := bytes.NewBuffer(u1.Bytes())
	binary.Read(buf, binary.BigEndian, &uuId)
	return
}

func Uuid(randomType string) string {
	return randomType + strconv.FormatUint(UuIdInt64(), 10)
}
