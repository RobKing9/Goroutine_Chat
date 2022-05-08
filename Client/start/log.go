package start

import "fmt"

/**
 * 注册程序
 **/
func clientReg(){
	for {
		var username,pwd,pwdCheck string
		fmt.Println("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pwd)
		fmt.Println("请确认密码：")
		fmt.Scanln(&pwdCheck)
		if pwdCheck != pwd {
			fmt.Println("两次输入密码不一致，请重新注册！")
		} else {
			sentData <- ("Reg-" + username + "-" + pwd)        // 发送用户名和密码
			res := <- recvData                                  // 等待登录结果
			if res == "success" {
				fmt.Println("注册成功，请登录")
				break
			}else{
				fmt.Println(res)
			}
		}
	}
	clientLog()      // 转至登录
}


/**
 * 登录程序
 **/
func clientLog(){

	LOOP: for {
		var username,password string
		fmt.Println("请输入用户名：")
		fmt.Scanln(&username)
		if username == "" {
			goto LOOP
		}
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)

		sentData <- "Log-" + username + "-" + password

		res :=<- recvData

		if res == "success" {

			fmt.Println("登录成功，你已进入聊天室！")

			uname = username    // 登录成功保存用户名

			go showMassage()    // 启动聊天内容显示线程

			break

		}else{

			fmt.Println(res)
		}
	}

	inputMessage()             // 启动聊天内容输入

}
