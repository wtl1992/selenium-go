package lang

import (
	"fmt"
	"hash/crc32"
	"log"
	"strconv"
)

type Object struct {
}

func (obj *Object) HashCode() uint32 {
	memAddress := fmt.Sprintf("%p", obj)

	memAddress = memAddress[2:]

	memInt, err := strconv.ParseUint(memAddress, 16, 64)

	if err != nil {
		log.Fatalln(err)
	}

	return crc32.ChecksumIEEE([]byte(strconv.FormatUint(memInt, 10)))
}
