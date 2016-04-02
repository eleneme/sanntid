package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os/exec"
	"time"
)

func spawnNewTerminal() {
	command := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run exercise6.go")
	_ = command.Run()
}

func main() {
	var master bool = false
	var counter uint64 = 0

	udpaddr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:30005")
	connection, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		fmt.Println("Housten, we have a problem!")
	}

	fmt.Println("backup running")
	UDPmsg := make([]byte, 8)

	for !(master) {
		connection.SetReadDeadline(time.Now().Add(time.Second * 2))

		n, _, err := connection.ReadFromUDP(UDPmsg)

		if err == nil {
			counter = binary.BigEndian.Uint64(UDPmsg[0:n])

		} else {
			master = true
		}

	}
	connection.Close()

	fmt.Println("Bitch, lay down, I am your master")
	spawnNewTerminal()
	connection, _ = net.DialUDP("udp", nil, udpaddr)

	for {

		fmt.Println(counter)
		counter++
		binary.BigEndian.PutUint64(UDPmsg, counter)
		_, _ = connection.Write(UDPmsg)

		time.Sleep(time.Second)
	}

}
