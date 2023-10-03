package main

import (
	"stocksocket/netws"
)

func main() {
	c := make(chan int)

	//gorilla.StartServer()
	netws.StartServer()

	<-c
}

//func messageHandler(message []byte) {
//	fmt.Println(string(message))
//}
