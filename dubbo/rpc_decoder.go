package dubbo

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Note: 简单版本
func Decode(res []byte) {
	readBytesLen := len(res)
	if readBytesLen < 16 {
		fmt.Println("need more data")
		return
	}
	var contentLen int32
	buf := bytes.NewReader(res[12:16])
	err := binary.Read(buf, binary.BigEndian, &contentLen)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("response length: ", contentLen)
	fmt.Println("response byte is: ", res[HEADER_LENGTH+2:readBytesLen-1])
}
