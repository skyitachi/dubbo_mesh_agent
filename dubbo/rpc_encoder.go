package dubbo

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

const HEADER_LENGTH = 16
const MAGIC = 0xdabb

var FLAG_REQUEST byte = 0x80
var FLAG_TWOWAY byte = 0x40
var FLAG_EVENT byte = 0x20

func Encode(invocation RpcInvocation) []byte {
	var header []byte
	var requestData []byte
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
	// TODO: 需要将 序列化到 bytes
	// dubbo version
	dubboByte, err := json.Marshal("2.0.1")
	if err != nil {

	}
	requestData = append(requestData, dubboByte...)

	pathByte, err := json.Marshal(invocation.Attachments["path"])
	if err != nil {

	}
	requestData = append(requestData, pathByte...)

	// NOTE: 没有该 key 就当 null 来处理
	versionByte, err := json.Marshal(invocation.Attachments["version"])

	if err != nil {

	}
	requestData = append(requestData, versionByte...)

	// write arguments
	requestData = append(requestData, invocation.Arguments...)

	// write attachments
	attachmentsBytes, err := json.Marshal(invocation.Attachments)
	requestData = append(requestData, attachmentsBytes...)

	// write length
	buf.Reset()
	err = binary.Write(buf, binary.BigEndian, len(requestData))
	if err != nil {

	}
	header = append(header, buf.Bytes()...)
	return append(header, requestData...)
}
