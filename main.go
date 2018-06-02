package main

import (
	"github.com/skyitachi/dubbo_mesh_agent/dubbo"
)

func main() {
	// http.HandleFunc("/", controller.HelloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	rpcClient := &dubbo.RpcClient{}
	rpcClient.Invoke(
		"",
	)
}
