package main

import (
	"fmt"
)

type Phone interface {//接口的定义
	call()
}

type NokiaPhone struct {//定义结构体
}

func (nokiaPhone NokiaPhone) call() {//定义接口函数
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {//定义结构体
}

func (iPhone IPhone) call() {//定义接口函数
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone//定义接口实例

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}