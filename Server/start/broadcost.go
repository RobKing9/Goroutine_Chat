package start

import (
	"fmt"
	"strings"
)

/**
* 服务器向客户端发送消息数据解析与封装操作
* @param client  客户端连接信息
* @param messages 数据通道中的数据
         群发数据格式：普通字符串
         单对单发送格式："To-" + uname(发送用户) + "-" + objUser（目标用户） +"-" +input
         用户列表："List-"+objUser（目标用户）+"-"+Listinfo
         命令:objUser + "-" + command
**/
func dataSent(conns *map[string]clientInfo,messages chan string){

	for{
		msg:= <- messages

		fmt.Println(msg)

		data := strings.Split(msg,"-")                         // 聊天数据分析:
		length := len(data)

		if length == 2 {                                       // 管理员单个用户发送控制命令

			(*conns)[data[0]].sentData <- data[1]

		} else if length == 3 {                                // 用户列表

			(*conns)[data[1]].sentData <- data[2]

		} else if length == 4 {                                // 向单个用户发送数据

			msg = data[1] + " say to you : " + data[3]

			(*conns)[data[2]].sentData <- msg

		} else {
			// 群发
			for _,value:= range *conns {
				value.sentData <- msg
			}
		}

	}
}

