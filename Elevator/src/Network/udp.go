package Network

import(
	"net"
	"fmt"
)

func connect(ip string) (connection *net.UDPConn){
	var address string
	if ip == ""{
		//Broadcast
		address = "129.241.187.255:" + PORT
	}else{
		address = ip + ":" + PORT
	}

	sendAddr, err := net.ResolveUDPAddr("udp", address)
	conn, err2 := net.DialUDP("udp", nil, sendAddr)

	if err != nil || err2 != nil{
		fmt.Println("Error connecting")
		return nil
	}

	if ip == ""{
		fmt.Println("Broadcast connection established")
		Broadcast.Conn = conn
	}else{
		fmt.Println("Connected to: ", ip)
		AppendConn(conn, ip)
	}

	return conn
}

func recieve(conn *net.UDPConn){
	//Recieve message from network
	recieved := make([]byte, 500)
	for; true; {
		_, _, _ = conn.ReadFromUDP(recieved)
		go recieveMessage(string(recieved))
	}
	conn.Close()
}

func send(s string, conn *net.UDPConn){
	b := []byte(s)
	_, _ = conn.Write(b)
}