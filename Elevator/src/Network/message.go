package network

import(
	"strings"
	"net"
	"time"
)
type Message struct{
	From string
	To string
	Msg string
	MsgType MessageType
}

type MessageType int
const(
	MasterAlive = iota 
	SlaveAlive
	MessageAck
	OrderCost //from slave
	NewOrder //from slave
	CostCalculated //from master
	SendOrder //from master to slave that is to execute the order
	BackUP //master sends backup to slaves
	NewMaster 
	OrderExecuted //slave informs that order is executed 
	
)

type sentMessage struct{
	To *net.UDPConn
	message *Message
	timeSent time.time
}

var sentMessages []*sentMessage

func SendMessage(text string, conn *net.UDPConn, ack bool){
	msg := new(Message)
	msg.From = IP 
	msg.Message = text
	messageString := msg.From + "+" + msg.Message

	if (ack){
		m.To = MasterConn.IP
		addMessage(m, conn)
	}

	send(messageString, conn)

}
//Funker dette?
func SendMessageType(m *Message, conn *net.UDPConn, ack bool){
	msg := new(Message)
	msg.From = IP 
	msg.MessageType = m.MessageType
	messageString := msg.From + "+" + msg.MessageType

	if (ack){
		m.To = MasterConn.IP
		addMessage(m, conn)
	}

	send(messageString, conn)

}

func recieveMessage(m string){
	msg := new(Message)
	text := string.Split(m, "+")
	msg.From = text[0]
	msg.Message = text[1]

	if msg.From != IP{
		handleMessage(msg)
	}
}

func printMessage(msg *Message){
	fmt.Println("")
	fmt.Prinln("From: ", msg.From)
	fmt.Println("Message: ", msg.Message)
	fmt.Println("")
}

func addMessage(message *Message, conn *net.UDPConn){
	var temp sentMessage
	temp.messageSent = message
	temp.timeSent = time.Now()
	temp.To = conn
	sentMessages = append(sentMessages, &temp) 
}

func messageAcknowledge(message *Message){
	for msg := 0; msg < len(sentMessages); msg++{
		if sentMessages[i].messageSent.Message == message.Message[2:6]{
			removeMessage(msg)
		}
	}
}

func removeMessage(msg int){
	sentMessage = append(sentMessages[:msg], sentMessages[msg+1]...)
}

func updateMessage(){
	for; true;{
		for i := 0; i < len(sentMessages); i++{
			if (time.Since(sentMessages[i].timeSent) > 2000*time.Millisecond){
				SendMessage(SendMessages[i].messageSent.Message, sentMessages[i].To, true)
				removeMessage(i)
			}
		}

		fmt.Println(len(sentMessages))
		time.Sleep(1000*time.Millisecond)
	}
}