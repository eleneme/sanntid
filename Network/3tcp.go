package main

import(
	"fmt"
	"net"
	"time"
	"bufio"
	"os"
	"log"
)

func errorCheck(err error){
	if err!= nil{
		log.Fatal("Error:", err.Error())
	}
}

func tcpSend(conn net.Conn, msg string){
	conn.Write([]byte(msg + string('\x00')))
}

func tcpReceive(conn net.Conn) string {
	msg, _ :=bufio.NewReader(conn).ReadString(byte('\x00'))
	return msg
}

func main(){
	serverAddr := "129.241.187.23:33546"
	fmt.Println("Launching TCP server")
	input := bufio.NewReader(os.Stdin)

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	errorCheck(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	errorCheck(err)

	fmt.Println(*tcpAddr)

	msgRcpt := tcpReceive(conn)
	fmt.Println(msgRcpt)

	defer conn.Close()

	for{

		fmt.Print("Enter message to send to sever: ")
		msgSend, _ := input.ReadString('\n')
		tcpSend(conn, msgSend)

		msgRcpt := tcpReceive(conn)

		fmt.Print("Message receiced: ", string(msgRcpt))

		time.Sleep(1)
	}

}