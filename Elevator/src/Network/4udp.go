package main

import (
	"fmt"
	"net"
)

//Sets up UDP functions for broadcasting and listening

func errorCheck(err error) {
	if err != nil {
		fmt.Println("Error", err)
	}
}
func UDPOpenBroadcastSocket(IP string) *net.UDPConn {
	UDPAddress, err := net.ResolveUDPAddr("udp", IP)
	errorCheck(err)

	connection, err := net.DialUDP("udp", nil, UDPAddress)
	errorCheck(err)

	return connection
}

func UDPOpenListenSocket(port string) *net.UDPConn {
	UDPAddress, err := net.ResolveUDPAddr("udp", port)
	errorCheck(err)

	connection, err := net.ListenUDP("udp", UDPAddress)
	errorCheck(err)

	return connection
}

func UDPBroadcast(conn *net.UDPConn, message string) {
	_, err := conn.Write([]byte(message))
	errorCheck(err)
}

func UDPListen(conn *net.UDPConn, chanListen <-chan config.NetworkMessage) {
	for {
		message := make([]byte, 1024)
		length, recieveAddress, _ := conn.ReadFromUDP(message)
		recievedMessage := config.NetworkMessage{recieveAddress: recieveAddress.IP.String(), data: message, length: length}
		chanListen <- recievedMessage
	}
}
