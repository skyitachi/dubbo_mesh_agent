package controller

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/skyitachi/dubbo_mesh_agent/dubbo"
)

func consume(interfaceName string, method string, parameterTypesString string, parameter string) {
	// TODO: 需要负载均衡
	endpoint := "http://localhost:30000"
	form := url.Values{
		"interfaceName":        {interfaceName},
		"method":               {method},
		"parameterTypesString": {parameterTypesString},
		"parameter":            {parameter},
	}
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(endpoint, "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Println("err is: ", err)
	}
	defer rsp.Body.Close()
	fmt.Println("call ok")
}

// HelloHandler Content-Type: application/x-www-form-urlencoded
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	interfaceName := r.Form.Get("interface")
	method := r.Form.Get("method")
	parameterTypesString := r.Form.Get("parameterTypesString")
	parameter := r.Form.Get("parameter")

	fmt.Fprintf(w, "hello world")
	invocation := dubbo.NewRpcInvocation(method, parameterTypesString)
	fmt.Println(dubbo.Encode(invocation))
	log.Print("interfaceName: ", interfaceName, method, parameterTypesString, parameter)
}
