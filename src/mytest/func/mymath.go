package main

import (
	"errors"
)

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-geative numbers.")
		return
	}
	return a + b, nil
}

//必须是main包下的main方法才可以运行
// func main() {
// 	res, err := Add(1, 2)
// 	if err != nil {
// 		fmt.Println("err")
// 	}
// 	fmt.Println("res=", res)
// }
