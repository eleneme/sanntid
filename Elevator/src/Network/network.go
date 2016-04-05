package Network

import (
	//"../config"
	"fmt"
	"net"
	"time"
	"strings"
	"strconv"
)

type Connection struct{
	IP string
	Conn *net.UDPConn
	LastSignal time.Time
}


var chNewMessage = make(chan *Message, 200)


var IP string
var PORT = "30005"
var Connected []Connection
var MasterConn Connection
var Broadcast Connection
var Master = false
var Alive bool

func NetworkInit() {
	//Retrive local IP address
	adr, _ := net.InterfaceAddrs()
	ip := strings.Split(adr[1].String(), "/")
	IP = ip[0]

	//Connect as listener
	recAddr, _ := net.ResolveUDPAddr("udp",":" + PORT)
	recConn, _ := net.ListenUDP("udp", recAddr)

	//Start listening-thread
	go recieve(recConn)

	//Check for master on network

	MasterConn.LastSignal = time.Now()
	var temp time.Time
	temp = time.Now()

	go timeout()

	for ; (time.Since(temp) < 350*time.Millisecond) && (Master == false);{

	}

	fmt.Println("Master: ", Master)

	//Establish broadcast-connection
	conn := connect("") 

	Alive = true
	go alive(conn)
}

//funker dette?
func alive(conn *net.UDPConn){
	//sends alive message
	for; Alive; {
		if  Master{
			//msg := new(Message)
			//msg.MsgType = MasterAlive
			SendMessage("MasterAlive", conn, false) 
			time.Sleep(100*time.Millisecond)
		}else{
			//msg := new(Message)
			//msg.MsgType = SlaveAlive
			SendMessage("SlaveAlive", conn, false) 
		}
	}
}

func handleMessage(msg *Message){
	order := msg.MsgType

	if (msg.From == IP){
		return
	}

	switch(order){
	case 0: //MasterAlive
		//Alive-signal from master
		MasterConn.LastSignal = time.Now()

		if Master{
			//Another master in Network  -> switch to slave
			Master = false
			fmt.Println("New master on network")
			MasterConn.IP = msg.From
			MasterConn.LastSignal = time.Now()
		}

	case 1: //SlaveAlive
		found := false

		for i := 0; i < len(Connected); i++{
			if msg.From == Connected[i].IP{
				Connected[i].LastSignal = time.Now()
				found = true
			}
		}

		if (!found) && (msg.From != MasterConn.IP){
			//new slave connexts
			connect(msg.From)
		}

	case 2: //MessageAck
		//Previous message acknowledged
		messageAcknowledge(msg)

	default:
		chNewMessage <- msg //sends msg to channelHandler NewMessage

	}

}

func timeout(){
	for; true; {

		if (time.Since(MasterConn.LastSignal) > 900*time.Millisecond) && (Master == false){
			//No master on network
			go FindNewMaster() //whosmaster

			//Remove messages waiting for acknowledgement from master

			for i := 0; i < len(sentMessages); i++{
				removeMessage(i)
			}

			time.Sleep(500*time.Millisecond)
		}

		for i := 0; i < len(Connected); i++{
			if time.Since(Connected[i].LastSignal) > 1200*time.Millisecond{
				//slave timed out

				fmt.Println("Slave timed out: ", Connected[i].IP)

				RemoveConn(i)

				for i := 0; i < len(Connected); i++{
					fmt.Println("Connected[",i,"]: ", Connected[i].IP)
				}

				if Master {
					//Distribute backup

					var temp Message
					temp.From = IP
					temp.MsgType = NewMaster
					chNewMessage <- &temp //NewMessage
				}
			}
		}

		time.Sleep(50*time.Millisecond)
	}

	fmt.Println("Timeout done")
}

func AppendConn(conn *net.UDPConn, ip string){
	var temp Connection
	temp.IP = ip
	temp.Conn = conn
	temp.LastSignal = time.Now()

	Connected = append(Connected, temp)
}

func RemoveConn(index int){
	Connected = append(Connected[:index], Connected[index+1:]...)
}

func FindConn(ip string) (*net.UDPConn){
	for i := 0; i < len(Connected); i++{
		if Connected[i].IP == ip{
			return Connected[i].Conn
		}
	}
	return nil
}

func FindNewMaster(){
	//finds new master with the lowest ip
	fmt.Println("Master timeout: finding new master")
	me := true
	number := len(IP)
	own, _ := strconv.Atoi(string(IP[number-2:number]))

	for i := 0; i < len(Connected); i++{
		fmt.Println("Connected[",i,"]: ", Connected[i].IP)
	}

	//compares own ip to lowest ip
	for i := 0; i < len(Connected); i++{
		other, _ := strconv.Atoi(string(Connected[i].IP[number-2:number]))

		if other < own{
			me = false
		}
	}

	if(me == true){
		Master = true

		//Distribute backup
		var temp Message
		temp.From = IP
		temp.MsgType = NewMaster
		chNewMessage <- &temp //NewMessage
	}

	MasterConn.IP = ""
	MasterConn.Conn = nil
}