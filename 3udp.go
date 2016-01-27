package main

import (
	"fmt"
	"net"
	"time"
)

const localAddr = "129.241.187.143"
const broadcastAddr = "129.241.187.255"
const receivePort = "30000"
const sendPort = "20005"

func errorCheck(err error){
	if err!= nil{
		fmt.Println("Error", err)
	}
}



func main() {
	serverAddr, err := net.ResolveUDPAddr ("udp", ":30000")
	errorCheck(err)
 
	buffer := make([]byte,1024)

    listenServer, err := net.ListenUDP("udp", serverAddr)
    errorCheck(err)

    defer listenServer.Close()

    	bytes, addr, err := listenServer.ReadFromUDP(buffer)
    	fmt.Println("Message recieved :", string(buffer[0:bytes]), "from address", addr)
    	errorCheck(err)

    serverAddrPort := "129.241.187.23:20005"

    localAddr, err := net.ResolveUDPAddr("udp", serverAddrPort)
    errorCheck(err)


    connectServer, err := net.DialUDP("udp", nil, localAddr)
    errorCheck(err)

   defer connectServer.Close()


   message := "Hello its me"
   for{
	   _, err = connectServer.Write([]byte(message + "\x00"))
	   errorCheck(err)

	   //buffer := make([]byte,1024)

	   numBytes, _ , _ := connectServer.ReadFromUDP(buffer)
	   fmt.Println(buffer[:numBytes])
	   time.Sleep(1)
	}


}