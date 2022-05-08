package start

import (
	"fmt"
	"net"
	"os"
	"time"
)

func Start() {
	// 创建TCP连接
	conn := createTCP()
	// 启动数据接收线程
	go recv(recvData, conn)
	// 启动数据发送线程
	go sent(sentData, conn)
LOOP:
	{
		var input string
		fmt.Scanln(&input)
		if input == "1" {
			clientLog() // 登录
		} else if input == "2" {
			clientReg() // 注册
		} else {
			fmt.Println("输入有误，请重新输入！")
			goto LOOP
		}

	}
}

/**
 * 创建TCP连接
 * @param  tcpAddr  TCP地址格式
 * return net.Conn
 **/
func createTCP() net.Conn {
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	//checkError(err,"ResolveTCPAddr")
	conn, err := net.Dial("tcp", "124.223.88.139:8080")
	checkError(err, "DialTCP")
	fmt.Println("正在连接中...")
	time.Sleep(2 * time.Second)
	fmt.Println("连接服务器成功！")
	fmt.Println("请输入将要进行的操作：1、登录  2、注册")
	return conn
}

/**
 * 错误检查
 * @param err   错误类型数据
 * @param info  错误信息提示内容
 **/
func checkError(err error, info string) {

	if err != nil {

		fmt.Println(info + "  " + err.Error())

		os.Exit(1)

	}

}
