package main
//函数方法
import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {//定义一个结构体
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}