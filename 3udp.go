package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const localAddr = "129.241.187.143"
const broadcastAddr = "129.241.187.255"
const receivePort = "20005"
const sendPort = "30005"

func errorCheck(err error) {
	if err != nil {
		fmt.Println("Error", err)
	}
}

func UDPListen() {

	serverAddr, err := net.ResolveUDPAddr("udp", ":20005")
	errorCheck(err)

	listenServer, err := net.ListenUDP("udp", serverAddr)
	errorCheck(err)
	buffer := make([]byte, 1024)
	numBytes, _, _ := listenServer.ReadFromUDP(buffer)
	fmt.Println(string(buffer[:numBytes]))

}
func UDPSend() {
	serverAddrPort := "129.241.187.20:54321"
	localAddr, err := net.ResolveUDPAddr("udp", serverAddrPort)
	errorCheck(err)

	connectServer, err := net.DialUDP("udp", nil, localAddr)
	errorCheck(err)
	input := bufio.NewReader(os.Stdin)
	msgSend, _ := input.ReadString('\n')

	_, err = connectServer.Write([]byte(msgSend + "\x00"))
}

func main() {
	/*serverAddr, err := net.ResolveUDPAddr ("udp", ":30000")
		errorCheck(err)

		buffer := make([]byte,1024)

		listenServer, err := net.ListenUDP("udp", serverAddr)
		errorCheck(err)

		defer listenServer.Close()

	    	bytes, addr, err := listenServer.ReadFromUDP(buffer)
	    	fmt.Println("Message recieved :", string(buffer[0:bytes]), "from address", addr)
	    	errorCheck(err)

	*/
	/*serverAddrPort := "129.241.187.20:54321"

	localAddr, err := net.ResolveUDPAddr("udp", serverAddrPort)
	errorCheck(err)

	connectServer, err := net.DialUDP("udp", nil, localAddr)
	errorCheck(err)

	serverAddr, err := net.ResolveUDPAddr("udp", ":20005")
	errorCheck(err)

	listenServer, err := net.ListenUDP("udp", serverAddr)
	errorCheck(err)

	defer listenServer.Close()

	defer connectServer.Close()

	message := "Hello can you hear ne"

	_, err = connectServer.Write([]byte(message + "\x00"))
	errorCheck(err)
	*/

	for {

		go UDPSend()
		go UDPListen()
	}

}
