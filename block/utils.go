package block

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(intNum int64) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, intNum)
	return buf.Bytes()
}
