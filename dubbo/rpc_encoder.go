package dubbo

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
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

	// 第四个字节
	header = append(header, 0x00)

	// write invocation id
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, invocation.ID)
	if err != nil {
		log.Printf("encode binary error %s\n", err.Error())
	}
	fmt.Println("invocation id bytes length: ", len(buf.Bytes()))
	header = append(header, buf.Bytes()...)
	// TODO: 需要将 序列化到 bytes
	// dubbo version
	dubboByte, err := json.Marshal("2.0.1")
	if err != nil {

	}
	requestData = append(requestData, append(dubboByte, 0x0a)...)

	pathByte, err := json.Marshal(invocation.Attachments["path"])
	if err != nil {

	}
	requestData = append(requestData, append(pathByte, 0x0a)...)

	// NOTE: 没有该 key 就当 null 来处理
	// 默认加上换行符
	v, ok := invocation.Attachments["version"]
	var versionByte []byte
	if !ok {
		v = "null"
		versionByte = []byte{0x6e, 0x75, 0x6c, 0x6c}
	} else {
		versionByte, err = json.Marshal(v)
		if err != nil {

		}
	}

	requestData = append(requestData, append(versionByte, 0x0a)...)
	// write method name
	methodNameByte, err := json.Marshal(invocation.Method)
	if err != nil {

	}
	requestData = append(requestData, append(methodNameByte, 0x0a)...)
	// write parameterTypes

	parameterTypesBytes, err := json.Marshal(invocation.ParameterTypes)
	if err != nil {

	}
	requestData = append(requestData, append(parameterTypesBytes, 0x0a)...)

	// write arguments
	requestData = append(requestData, append(invocation.Arguments, 0x0a)...)

	// write attachments
	attachmentsBytes, err := json.Marshal(invocation.Attachments)
	requestData = append(requestData, append(attachmentsBytes, 0x0a)...)

	// write length
	lenBuf := new(bytes.Buffer)
	fmt.Println("data length: ", len(requestData))
	err = binary.Write(lenBuf, binary.BigEndian, int32(len(requestData)))
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("buffer: ", len(lenBuf.Bytes()))
	header = append(header, lenBuf.Bytes()...)
	return append(header, requestData...)
}
