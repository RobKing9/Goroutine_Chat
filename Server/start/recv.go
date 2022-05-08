package start

import "net"

// 声明连接管理Channel
var connection  = make(chan net.Conn)
/**
 * 数据接收
 * @param  recvData  接收数据Channel
 * @param  conn      TCP连接对象指针
 **/
func recv(recvData chan string,conn net.Conn){

	for{

		buf := make([]byte,1024)

		n,err := conn.Read(buf)

		if err != nil {

			connection <- conn

			return
		}

		recvData <- string(buf[:n])

	}

}

