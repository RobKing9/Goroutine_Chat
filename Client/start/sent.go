package start

import (
	"fmt"
	"net"
	"os"
)

// 发送数据Channel
var sentData = make(chan string)
// 用户名
var uname string

/**
 * 数据发送
 * @param  sentData  接收数据Channel
 * @param  conn      TCP连接对象
 **/
func sent(sentData chan string,conn net.Conn){

	for{
		data := <-sentData
		_,err := conn.Write([]byte(data))
		checkError(err,"Connection")
	}
}

/**
 * 聊天数据输入
 **/
func inputMessage(){
	var input string
	for {
		fmt.Scanln(&input)
		switch input {
			// 用户退出
			case "/quit":  fmt.Println("退出聊天室，欢迎下次使用！")
				os.Exit(0)
			// 默认群发
			default:   if len(input) != 0 {
				input = uname + " : " + input
				}
		}
		// 发送数据
		if len(input) != 0 {
				sentData <- input
			input = ""
		}
	}
}

/**
 * 聊天内容显示
 **/
func showMassage(){
	for{
		message :=<- recvData
			fmt.Println(message)
	}
}

