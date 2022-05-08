package start

import "net"

// 接受数据Channel
var recvData = make(chan string)

/**
 * 数据接收
 * @param  recvData  接收数据Channel
 * @param  conn      TCP连接对象指针
 **/
func recv(recvData chan string,conn net.Conn){

	for{

		buf := make([]byte,1024)

		n,err := conn.Read(buf)
	
		checkError(err,"Connection")

		recvData <- string(buf[:n])

	}

}
