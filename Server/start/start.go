package start

import (
	"fmt"
	"net"
	"os"
)
// 声明客户端的结构体
type clientInfo struct{
	// 客户端的TCP连接对象
	conn net.Conn
	// 服务器向客户端发送数据通道
	sentData chan string
}

// 声明成功登录之后的连接对象map
var conns  = make(map[string] clientInfo)
// 声明消息channel
var messages = make(chan string)
/**
 * 服务端启动程序
 * @param  port  设置监听端口
 **/
func StartServer(){
	//连接
	l := createTCP()
	fmt.Println("服务端启动成功，正在监听端口！")
	// 启动服务器广播线程
	go dataSent(&conns,messages)
	//go intoManager()              // 启动管理模块
	//go connManager(connection)    // 启动连接管理线程
	for  {
		conn,err := l.Accept()
		if checkError(err,"Accept") == false {
			continue
		}
		fmt.Println("客户端:",conn.RemoteAddr().String(),"连接服务器成功！")
		var recvData = make(chan string)
		var sentData = make(chan string)
		// 开启对客户端的接受数据线程
		go recv(recvData,conn)
		// 开启对客户端的发送数据线程
		go sent(sentData,conn)
		// 用户资格认证
		go userAuth(&conn,recvData,sentData)
	}

}
/**
 * 设置TCP连接
 * @param  tcpAddr  TCP地址格式
 * return net.TCPListener
 **/
func createTCP() net.Listener{
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	//errorExit(err,"ResolveTCPAddr")
	//l,err := net.ListenTCP("tcp",tcpAddr)
	l,err := net.Listen("tcp",":8080")
	errorExit(err,"DialTCP")
	return l
}
/**
 * 错误显示并退出
 * @param err   错误类型数据
 * @param info  错误信息提示内容
 **/
func errorExit(err error,info string){
	if err != nil {
		fmt.Println(info+"  " + err.Error())
		os.Exit(1)
	}
}
/**
 * 错误检查
 * @param err   错误类型数据
 * @param info  错误信息提示内容
 **/
func checkError(err error,info string) bool{
	if err != nil {
		fmt.Println(info+"  " + err.Error())
		return false
	}
	return true
}
