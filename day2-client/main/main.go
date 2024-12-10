package main

import(
	"fmt"
	"gorpc"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string){
	l,err:=net.Listen("tcp",":0")
	if err!=nil{
		log.Fatal("network error:",err)
	}
	log.Println("start rpc server on",l.Addr())
	addr<-l.Addr().String()
	gorpc.Accept()
}

