package main

import (
	"github.com/skyitachi/dubbo_mesh_agent/dubbo"
)

func main() {
	// http.HandleFunc("/", controller.HelloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	done := make(chan int)
	rpcClient := &dubbo.RpcClient{}
	go func() {
		rpcClient.Invoke(
			"com.alibaba.dubbo.performance.demo.provider.IHelloService",
			"hash",
			"Ljava/lang/String;",
			"1",
		)

	}()
	<-done
}
