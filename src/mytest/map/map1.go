package main

import (
	"fmt"
	// mymath "mytest/func"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "width id 1234.")

	} else {
		fmt.Println("Did not find person with id 1234")
	}
	// var myMap map[string]string = nil
	// delete(myMap, "1234")
	// fmt.Println(testIfReturn(6))
	// testFor()
	// testFor1()
	testForBreak()
	// res, err := mymath.Add(1, 2)
	// if err != nil {
	// 	fmt.Println("err")
	// }
	// fmt.Println("res=", res)
}
func testIfReturn(x int) int {
	if x < 5 {
		return 0
	}
	if x == 6 {
		return 6
	} else {
		return x
	}

}
func testFor() {
	sum := 0
	for index := 0; index < 10; index++ {
		fmt.Println("sum1=", sum)
	}
}
func testFor1() {
	sum := 0
	for {
		sum++
		fmt.Println("sum2=", sum)
		if sum >= 10 {
			break
		}
	}
}
func testForBreak() {
	i, j := 0, 0
	fmt.Printf("%+v\n", 11)
	for {
		j++
		if j == 5 {
			break
		}
	ScondLoop:
		for {
			i++
			if i == 10 {
				break
			}
			if i > 5 {
				fmt.Printf("break second loop ,j=%v, i=%v\n", j, i)
				goto JLoop
				break ScondLoop
			}
			if i == 8 {
				fmt.Printf("goto JLoop ,j=%v, i=%v\n", j, i)
				goto JLoop
			}

			fmt.Printf("j=%f, i=%v\n", j, i)
		}
	}
JLoop:
	{
		fmt.Println("this is JLoop")
	}

}
