package dubbo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

type RpcClient struct {
}

func (rpcClient *RpcClient) Invoke(
	interfaceName string,
	method string,
	parameterTypesString string,
	parameter string,
) {
	// TODO: get remote address

	// TODO: serialize data
	done := make(chan int)
	invocation := NewRpcInvocation(method, parameterTypesString)
	invocation.Attachments["path"] = interfaceName
	paramBytes, err := json.Marshal(parameter)
	if err != nil {
		log.Printf("marshal json data %s error %s\n", parameter, err.Error())
		return
	}
	invocation.Arguments = paramBytes
	encoded := Encode(invocation)
	// TODO: just test code
	conn, err := net.Dial("tcp", "localhost:20880")
	fmt.Println("local address: ", conn.LocalAddr().String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("encoded: ")
	fmt.Printf("%x\n", encoded)
	writed, err := conn.Write(encoded)

	if writed == len(encoded) {
		log.Println("write data to request right")
	}
	var resp []byte
	var buf = make([]byte, 1024)
	for {
		respL, err := conn.Read(buf)
		if err != nil {
			fmt.Println("in the error: ", err)
			if err != io.EOF {
				fmt.Println("read error: ", err)
			}
			log.Fatal(err)
			break
		}
		if respL != 0 {
			// Note: 先简单处理
			fmt.Println("in the response: ", buf[:respL])
			resp = append(resp, buf[:respL]...)
			Decode(buf[:respL])
			break
		}
	}
	// for _, b := range resp {
	// 	fmt.Printf("% 20x", b)
	// }
	// fmt.Printf("response bytes: % x\n", resp)

	// defer conn.Close()
	<-done
}
