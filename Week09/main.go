package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	messageChan chan []byte
	conn        net.Conn
}

var IpList map[string]net.Conn
var CloseChan chan int

func main() {
	CloseChan = make(chan int, 1)
	IpList = map[string]net.Conn{}
	tcp, err := net.Listen("tcp", ":8200")
	if err != nil {
		log.Fatal("err is %v", err)
	}
	for {
		m := Message{}
		m.conn, err = tcp.Accept()
		m.messageChan = make(chan []byte, 1)
		//连接过来的客户端ip都记录起来
		log.Println("客户端连接来了", m.conn.RemoteAddr())
		IpList[m.conn.RemoteAddr().String()] = m.conn
		if err != nil {
			fmt.Println("error is read")
		}
		fmt.Println(IpList)
		// 启动读写的channel
		go write(&m)
		go read(&m)

	}
}
func CloseSingle() {
	CloseChan <- 1
}

func write(m *Message) {
	defer m.conn.Close()
	for {
		select {
		case data := <-m.messageChan:
			//消息成功失败状态
			msgStatus := "1"
			//给iplist里边的指定ip发送消息
			othrconn := IpList[string(data)]
			n, err := othrconn.Write(data)
			if err != nil {
				msgStatus = "0"
			}
			fmt.Println("消息发送成功 消息字节数-", n)
			//给本机发送回执
		E:
			n, err = m.conn.Write([]byte(msgStatus))
			if err != nil {
				goto E
			}
		case <-CloseChan:
			//当前客户端下线
			delete(IpList, m.conn.RemoteAddr().String())
			//跳出循环
			break
		default:
		}
	}
}
func read(m *Message) {
	// 读取32个字节
	for {
		tmpByte := make([]byte, 1024*1024)
		n, err := m.conn.Read(tmpByte)
		if n == 0 || err != nil {
			CloseSingle()
			break
		}
		m.messageChan <- tmpByte[:n]
	}

}

// 处理函数
//func process(conn net.Conn) {
//	defer conn.Close() // 关闭连接
//	for {
//		reader := bufio.NewReader(conn)
//		var buf [128]byte
//		n, err := reader.Read(buf[:]) // 读取数据
//		if err != nil {
//			fmt.Println("read from client failed, err:", err)
//			break
//		}
//		recvStr := string(buf[:n])
//		fmt.Println("收到client端发来的数据：", recvStr)
//		conn.Write([]byte(recvStr)) // 发送数据
//	}
//}
//
//func main() {
//	listen, err := net.Listen("tcp", ":8200")
//	if err != nil {
//		fmt.Println("listen failed, err:", err)
//		return
//	}
//	for {
//		conn, err := listen.Accept() // 建立连接
//		if err != nil {
//			fmt.Println("accept failed, err:", err)
//			continue
//		}
//		go process(conn) // 启动一个goroutine处理连接
//	}
//}
