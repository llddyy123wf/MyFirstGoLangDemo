package main

import (
	"fmt"
	"time"
)

func main() {
	testMultyPara(1, 2, 3)
	//测试任意类型多参数传递
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float64 = 1.234
	v5 := 5
	testMultyTypeMultyPara(v1, v2, v3, v4, v5)
	//测试匿名函数,赋给变量直接调用
	f := func(a, b int, c float64) bool {
		return a*b < int(c)
	}
	fmt.Println("匿名函数结果：", f(1, 2, 3.5))
	//测试匿名函数，也可以申明扣直接调用
	func(ch int) {
		fmt.Println(ch)
	}(11)
	//使用go关键字代表使用了多线程，多线程情况下，因为程序先执行主线程，如果不给休眠时间，子线程没时间执行就退出了
	go func(ch int) {
		fmt.Println(ch)
	}(22)
	time.Sleep(1)
	//测试闭包(闭包可以保证闭包内的变量的安全性)
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}

//测试多参数
func testMultyPara(args ...int) {
	for _, arg := range args {
		fmt.Printf("arg=%v\n", arg)
	}
}

//测试任意类型的多参数
func testMultyTypeMultyPara(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int type")
		case string:
			fmt.Println(arg, "is a string type")
		case int64:
			fmt.Println(arg, "is an int64 type")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

// func (e *PathError) Error() string {
// 	return e.Op + " " + e.Path + ": " + e.Err.Error()
// }
// func Stat(name string) (fi FileInfo, err error) {
// 	var stat syscall.Stat_t
// 	err = syscall.Stat(name, &stat)
// 	if err != nil {
// 		return nill, &PathError{"stat", name, err}
// 	}
// 	return fileInfoFromSta(&stat, name), nil
// }
