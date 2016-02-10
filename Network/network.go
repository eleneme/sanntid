package main

import (
	"config"
	"fmt"
	"net"
)

const (
	UDPPORT       string = ":30005"
	TCPPORT       string = ":20005"
	BROADCAST     string = "129.241.187.255"
	IPBASE        string = "129.241.187."
	BUFLEN        int    = 1024
	IPLEN         int    = 12
	LOCALIPSERVER string = "23"
)

func networkInit() {
	chTCPTransmit := make(chan config.NetworkMessage)
	chTCPRecieve := make(chan config.NetworkMessage, 1024)
	chUDPTransmit := make(chan config.NetworkMessage)
	chUDPRecieve := make(chan config.NetworkMessage, 1024)

}
