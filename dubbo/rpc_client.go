package dubbo

import (
	"encoding/json"
	"log"
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
}
