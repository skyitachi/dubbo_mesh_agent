package dubbo

import (
	"encoding/json"
	"fmt"
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
	if err != nil {
		log.Fatal(err)
	}
	writed, err := conn.Write(encoded)
	if writed == len(encoded) {
		log.Println("write data to request right")
	}
	var resp []byte
	_, err = conn.Read(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response bytes: ", resp)

	defer conn.Close()
	fmt.Println(encoded)
}
