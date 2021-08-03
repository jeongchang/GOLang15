package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨을 켠다")
	} else if temp <= 3 {
		fmt.Println("turn on Heater")
	} else if temp < 18 {
		fmt.Println("go out")
	} else {
		fmt.Println("so hot")
	}
}
