package trydefer

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

//写一个文件
func writeFile(fileName string){
	file, e := os.Create(fileName)
	if nil != e {
		log.Println("异常：",e.Error())
		return
	}
	//关闭io
	defer file.Close()

	//使用缓冲提高效率
	writer := bufio.NewWriter(file)
	// 保证释放 写入
	defer writer.Flush()

	fmt.Fprintln(writer,"Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。")
}

type User struct {
	name,pwd string
}

type LoginError struct {
	msg string
}

//自定义异常
func (err *LoginError) Error() string {
	return err.msg
}

func login(name string,pwd string) (*User,error) {
	if name != "go" {
		return nil, &LoginError{msg:"登录账号异常..."}
	}

	if pwd != "123456" {
		return nil, &LoginError{msg:"密码错误..."}
	}

	return &User{name:name,pwd:pwd}, nil
}

func tryRecover()  {
	//使用 defer 保证该匿名方法总被执行
	defer func() {
		r := recover()
		if err,ok := r.(error); ok{
			fmt.Println("recover 异常信息是 ",err)
		}else{
			panic("不是error类型的错误.")
		}
	}()

	//制造一个异常
	//panic("这是一个错误...")
	panic(errors.New("这是一个错误..."))
}

func DeferDemo()  {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error....")
	//return
	//fmt.Println(4)

	user, e := login("go", "123456")
	if nil != e {
		loginError,ok:= e.(*LoginError)
		if !ok {
			//未知的异常类型 不处理 抛出
			panic(e)
		}else{
			fmt.Println(loginError.Error())
		}
	}else{
		fmt.Println(user)
	}

	writeFile("demo.txt")
	tryRecover()

}