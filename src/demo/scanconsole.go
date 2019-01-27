/*
从控制台中读取文件
*/

package main

import "fmt"

var (
	FirstName, SecondName, ThirdName string
	i                                int
	f                                float32
	Input                            = "5.2/100/yin"
	format                           = "%f/%d/%s"
)

//1从控制台读取输入
func testscan() {
	fmt.Println("1从控制台读取输入")
	fmt.Printf("Please enter your name:")
	fmt.Scanln(&FirstName, &SecondName) //扫描来自标准输入的文本，只到遇到换行
	// fmt.Scanf("%s %s", &FirstName, &SecondName) //与sclanln类似，不过第一个参数用于格式化字符串
	fmt.Printf("Hi %s %s!\n", FirstName, SecondName)
	fmt.Sscanf(Input, format, &f, &i, &ThirdName) //Sscan是从字符串中读取
	fmt.Println("From the input ,we read :", f, i, ThirdName)
}

func main() {
	// testscan()
	// testBufio()
	testBufio2()
}
