package phoenix

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var Master int
var LastSignal time.Time

func Phoenix() {
	Master = 0
	LastSignal = time.Now()

	slave()
	fmt.Println("Exciting phoenix")
}

func connect()(*net.UDPConn) {
	sendAddr, _ := net.ResolveUDPAddr("udp", "localhost:30005")
	conn, _ := net.DialUDP("udp", nil, sendAddr)
	return conn
}

func connRecieve()(*net.UDPConn) {
	recAddr, _ := net.ResolveUDPAddr("udp", ":30005")
	recConn, _ := net.ListenUDP("udp", recAddr)
	return recConn
}

func recieve(conn *net.UDPConn) {
	received := make([]byte, 1)
	for Master == 0 {
		_, _, _ = conn.ReadFromUDP(received)
		LastSignal = time.Now()
	}
	conn.Close()
}

func send(conn *net.UDPConn) {
	a, _ := json.Marshal("a")
	_, _ = conn.Write(a)

}

func alive() {
	conn := connect()
	for {
		send(conn)
		time.Sleep(1000 * time.Millisecond)
	}
}

func slave() {
	recConn := connRecieve()
	go recieve(recConn)
	for {
		if time.Since(LastSignal) > 2000*time.Millisecond {
			Master = 1
			go alive()
			return
		}
	}
}
