package dubbo

import (
	"bytes"
	"encoding/binary"
	"log"
)

const HEADER_LENGTH = 16
const MAGIC = 0xdabb

var FLAG_REQUEST byte = 0x80
var FLAG_TWOWAY byte = 0x40
var FLAG_EVENT byte = 0x20

func Encode(invocation RpcInvocation) []byte {
	var header []byte
	header = append(header, 0xda)
	header = append(header, 0xbb)
	header = append(header, (FLAG_REQUEST | 6))

	if invocation.isEvent {
		header[2] |= FLAG_EVENT
	}
	if invocation.isTwoWay {
		header[2] |= FLAG_TWOWAY
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, invocation.ID)
	if err != nil {
		log.Printf("encode binary error %s\n", err.Error())
	}
	header = append(header, buf.Bytes()...)
	return header
}
