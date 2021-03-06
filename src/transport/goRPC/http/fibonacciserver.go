package main

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
	"net/http"
	"strconv"
	"shared/parameters"
	"apps/fibonacci/implrpc"
	"shared/net"
)

func main() {

	// create new instance of fibonacci
	fibonacci := new(implrpc.Fibonacci)

	// create new goRPC server
	server := rpc.NewServer()
	server.RegisterName("Fibonacci", fibonacci)

	// associate a http handler to server
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, e := net.Listen("tcp", netshared.ResolveHostIp()+":"+strconv.Itoa(parameters.FIBONACCI_PORT))
	if e != nil {
		log.Fatal("Server not started:", e)
	}

	// wait for calls
	fmt.Println("Server is running at "+netshared.ResolveHostIp()+" Port "+strconv.Itoa(parameters.FIBONACCI_PORT)+"\n")
	http.Serve(l, nil)
}
