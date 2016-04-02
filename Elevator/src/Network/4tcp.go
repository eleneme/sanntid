package main

import (
	"bufio"
	"config"
	"fmt"
	"net"
)

//Sets up TCP functions for transmitting and recieving
func errorCheck(err error) {
	if err != nil {
		fmt.Println("Error", err)
	}
}

func TCPConnectRequest(IP string) *net.TCPConn {
	TCPAddress, err := net.ResolveTCPAddr("tcp", IP)
	errorCheck(err)

	connection, err := net.DialTCP("tcp", nil, TCPAddress)
	errorCheck(err)

	return connection
}

func TCPTransmit(connection *net.TCPConn, channelTransmit <-chan config.NetworkMessage) {
	for {
		message <- channelTransmit
		append(message.data, byte('\x00'))
		connection.Write(message.data)
	}
}

func TCPReceive(connection *net.TCPConn, channelRecieve <-chan config.NetworkMessage) {
	for {
		message, _ := bufio.NewReader(connection).ReadByte(byte('\x00'))
		recievedMessage := config.NetworkMessage{recievedAddress: connection.RemoteAddr(), data: message, length: len(message)}
		channelRecieve <- receivedMessage
	}
}
