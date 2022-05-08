package start

import "net"

/**
 * 数据发送
 * @param  sentData  接收数据Channel
 * @param  conn      TCP连接对象
 **/
func sent(sentData chan string,conn net.Conn){

	for{

		data := <-sentData

		_,err := conn.Write([]byte(data))

		if err != nil {

			connection <- conn

			return
		}

	}
}
