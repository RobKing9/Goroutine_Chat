package start

/**
 * 用户登录处理
 * @param username  用户登录名
 * @param password  用户密码
 * return 返回处理状态以及对应信息
 **/
func logProcess(username,password string) (status bool,info string){

	if len(conns) == maxLog {
		return false,"当前登录人数已满,请稍后登录"
	}
	if user,ok := uData[username] ; !ok {
		return false,"用户名或密码错误!"
	} else {
		if user.Password == password {
			if _,ok := conns[username]; ok{
				return false,"用户已登录！"
			} else{
				return true,success
			}
		}else{
			return false,"用户名或密码错误!"
		}
	}
}


/**
 * 用户注册处理
 * @param username  用户登录名
 * @param password  用户密码
 * return 返回处理状态以及对应信息
 **/
func regProcess(username,password string) (status bool,info string){
	if len(uData) == maxReg {
		return false,"注册人数已满！"
	}
	if _,ok := uData[username] ; ok {
		return false,"用户已存在，请更换注册名!"
	} else {
		insertNewUser(dataFileName,username,password)
		return true,success
	}
}

