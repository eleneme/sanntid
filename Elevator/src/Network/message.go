package Network

import(
	"strings"
	"net"
	"time"
	"fmt"
)
type Message struct{
	From string
	To string
	Message string
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
	timeSent time.Time
}

var sentMessages []*sentMessage

func SendMessage(text string, conn *net.UDPConn, ack bool){
	msg := new(Message)
	msg.From = IP 
	msg.Message = text
	messageString := msg.From + "+" + msg.Message

	if ack{
		msg.To = MasterConn.IP
		addMessage(msg, conn)
	}

	send(messageString, conn)

}

/*//Funker dette?
func SendMessageType(m *Message, conn *net.UDPConn, ack bool){
	msg := new(Message)
	msg.From = IP 
	msg.MsgType = m.MsgType
	msg.Message = ""
	messageString := msg.From + "+" + msg.Message

	if ack{
		m.To = MasterConn.IP
		addMessage(m, conn)
	}

	send(messageString, conn)

}
*/
func recieveMessage(m string){
	msg := new(Message)
	text := strings.Split(m, "+")
	msg.From = text[0]
	msg.Message = text[1]

	if msg.From != IP{
		handleMessage(msg)
	}
}

func printMessage(msg *Message){
	fmt.Println("")
	fmt.Println("From: ", msg.From)
	fmt.Println("Message: ", msg.Message)
	fmt.Println("")
}

func addMessage(message *Message, conn *net.UDPConn){
	var temp sentMessage
	temp.message = message
	temp.timeSent = time.Now()
	temp.To = conn
	sentMessages = append(sentMessages, &temp) 
}

func messageAcknowledge(message *Message){
	for msg := 0; msg < len(sentMessages); msg++{
		if sentMessages[msg].message.Message == message.Message[2:6]{
			removeMessage(msg)
		}
	}
}

func removeMessage(msg int){
	sentMessages = append(sentMessages[:msg], sentMessages[msg+1:]...)
}

func updateMessage(){
	for; true;{
		for i := 0; i < len(sentMessages); i++{
			if (time.Since(sentMessages[i].timeSent) > 2000*time.Millisecond){
				SendMessage(sentMessages[i].message.Message, sentMessages[i].To, true)
				removeMessage(i)
			}
		}

		fmt.Println(len(sentMessages))
		time.Sleep(1000*time.Millisecond)
	}
}