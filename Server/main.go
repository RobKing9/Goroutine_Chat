/*
	通过go run main/main.go 8080(端口号) 启动
*/

package main

import "Server/start"

// 常量配置
const (
	dataFileName = "database.txt" // 用户注册数据文件保存名
	dataSize     = 5 * 1024       // 文件大小
	maxLog       = 30             // 最大同时登录人数
	maxReg       = 500            // 最大用户注册人数
	success      = "success"      // 登录注册返回给客户端成功标识
)

func main() {
	start.GetAllUser(dataFileName) // 用户登录、注册数据初始化
	start.StartServer()            // 启动客户端

}
