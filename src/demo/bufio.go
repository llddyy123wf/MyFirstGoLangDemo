package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	inputReader *bufio.Reader //inputReader是一个指向bufio.Reader的指针
	input       string
	err         error
)

//2.从缓冲区读取输入
func testBufio() {
	fmt.Println("2.从缓冲区读取输入")
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter some input: ")
	//读取器对象提供一个方法ReadString(delim byte),该方法从输入中读取内容，
	//直到碰到delim指定的字符，然后将读取到的内容连同delim字符一起入到缓冲区
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s", input)
	}
}

//3.从键盘读取输入
func testBufio2() {
	fmt.Println("3.从键盘读取输入开始")
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name:")
	input1, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading,exiting program.")
		return
	}
	fmt.Printf("Your name is %sseesee", input1)
	fmt.Println("strings.ContainsAny(input1):", strings.ContainsAny(input1, "\r"))
	//去掉换行符之后的串
	// newstr := strings.Replace(input1, "\r\n", "", -1)

	//写法1
	// switch input {
	// case "ldy", "lidy\n", "deyin\n":
	// 	fmt.Println("Welcome ", input)
	// default:
	// 	fmt.Println("You are not welcome.Goodbye.")
	// }
	//写法2
	switch input1 {
	case "ldy":
		fmt.Println("222")
	case "ldy\n":
		fmt.Println("22")
	case "lidy\n":
		//fallthrough
	case "dy\r\n":
		fmt.Println("Welcome ", input1)
		// default:
		// 	fmt.Println("input:", input1)
		// 	fmt.Println("You are not welcome.Goodbye.")
	}
	if input1 == "ldy\n" {
		fmt.Println("this is ldy")
	} else {
		fmt.Println("this is ldy-n")
	}
	// fmt.Println("mm：", strings.TrimRight("abcdaaaa", "abcd"))
	// fmt.Println("mm：", strings.TrimRight("abcdabcdaaaaeac", "abcd"))
	// fmt.Println("mm：", strings.TrimLeft("abcdabcdaaaaeac", "abcd"))
	// fmt.Println("mm：", strings.TrimRight("dbdcaaaaabcd", "cd"))
	// fmt.Println("sf：", strings.TrimPrefix("abcdabcdaaaaeac", "abcd"))

}

//3.从键盘读取输入
func testBufio21() {
	fmt.Println("3.从键盘读取输入开始")
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name:")
	input1, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading,exiting program.")
		return
	}
	fmt.Printf("Your name is %s", input1)
	fmt.Println("strings.ContainsAny(input1):", strings.ContainsAny(input1, "\n"))
	//去掉换行符之后的串
	newstr := strings.Replace(input1, "\n", "", -1)
	fmt.Println("1.我不理解为什么不加ln会变成如下,首先是顺序不对，其次是少了x位，后来发现是好像被覆盖了，不知是vscode的问题还是golang的replace的问题：")
	fmt.Print(newstr)
	fmt.Println("xxx")
	fmt.Print(newstr)
	fmt.Println("xx")
	fmt.Println("2.加ln会变成如下：")
	fmt.Println(newstr)
	fmt.Println("sss")
	fmt.Println("3.我不用replace了,直接输出,这种情况下没有问题,因为输入的内容中有一个换行符，所以会换行：")
	fmt.Print(input1)
	fmt.Println("ss")

}

type Base struct {
	Name string
}

func (base *Base) Foo() {

}
func (base *Base) Bar() {

}

type Foo struct {
	Name string
	Base
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
}
