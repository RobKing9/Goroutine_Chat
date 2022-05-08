package start

import (
	"encoding/json"
	"net"
	"os"
	"strings"
)
// 常量配置
const (
	dataFileName  = "database.txt"    // 用户注册数据文件保存名
	dataSize = 5 * 1024               // 文件大小
	maxLog = 30                       // 最大同时登录人数
	maxReg = 500                      // 最大用户注册人数
	success = "success"               // 登录注册返回给客户端成功标识
)

// 用户数据结构体
type userData struct {
	Password  string
	Level int
}
// 声明登录用户数据map
var uData = make(map[string]userData)
// ip地址与用户名队名
var ipToUname = make(map[string]string)
/*
 * 获取全部已注册用户的数据,用于登录注册时验证
 * @param filename  文件名
 * return  userData  用户数据map
 */
func GetAllUser(filename string) map[string]userData{
	buf := make([]byte,dataSize)
	udata := make(map[string]userData)
	file,err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	defer file.Close()
	errorExit(err,"getAllUser")
	n,_:= file.Read(buf)
	json.Unmarshal(buf[:n],&udata)
	return udata
}

/**
 * 添加新注册用户数据
 * @param filename  文件名
 * @param username  用户名
 * @param password  密码
 */
func insertNewUser(filename,username,password string){
	file,err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	defer file.Close()
	checkError(err,"insertNewUser")
	uData[username] = userData{password,1}
	data, _ := json.MarshalIndent(uData, "", "  ")
	file.WriteString(string(data))
}

/**
 * 用户进入聊天室资格认证
 * @param recvData,sentData 客户端连接接收和发送Channel
 * @param conn TCP连接对象
 **/
func userAuth(conn *net.Conn,recvData,sentData chan string){

	for {
		// 等待用户发送登录或注册数据
		data := strings.Split(<- recvData,"-")
		flag,username,password := data[0],data[1],data[2]

		if flag == "Reg"{

			_,info := regProcess(username,password)
			sentData <- info

		} else {

			status,info := logProcess(username,password)
			sentData <- info

			if status == true {
				messages <- ("用户:"+username+"进入聊天室")                          // 用户登录成功，发送系统通知
				conns[username] = clientInfo{*conn,sentData}                        // 记录用户登录状态信息
				ipToUname[(*conn).RemoteAddr().String()] = username                 // 记录IP与用户名对应信息
				messages <- "List-"+username+"-"+userList()                         // 向用户发送已在线用户列表
				go dataRec(&conns,username,recvData,messages)                       // 开启服务器对该客户端数据接收线程
				break
			}
		}
	}
}
/**
 * 获取用户列表
 * @param string   用户列表
 **/
func userList() string{

	var userList string = "当前在线用户列表："
	for user := range conns {

		userList += "\r\n" + user

	}
	return userList
}

/**
 * 用户登录之后将用户接受的数据放入公共channel
 * @param messages 数据通道
 **/
func dataRec(conns *map[string]clientInfo,username string,recvData,messages chan string){
	for{
		data := <- recvData
		if _,ok := (*conns)[username];!ok{
			return
		}
		if len(data) > 0{
			messages <- data
		}

	}

}
