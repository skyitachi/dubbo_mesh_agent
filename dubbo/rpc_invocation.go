package dubbo

var counter int64 = 1

type RpcInvocation struct {
	InterfaceName  string
	Method         string
	ParameterTypes string
	Arguments      []byte
	Attachments    map[string]string
	isTwoWay       bool
	isEvent        bool
	ID             int64
}

func NewRpcInvocation(method string, parameterTypes string) RpcInvocation {
	invocation := RpcInvocation{
		Method:         method,
		ParameterTypes: parameterTypes,
	}
	invocation.Attachments = make(map[string]string)
	invocation.isEvent = false
	// Note: 暂时先默认为 true
	invocation.isTwoWay = true
	invocation.ID = counter
	counter++
	return invocation
}
