package main

import (
	"fmt"
)

func main() {
	fmt.Println(Prompt("My Web Store in Golang"))
	fmt.Println(Out("Output: ", GreenText("Green"), BoldText("White"), RedBoldText("Red")))

	fmt.Println()
	fmt.Print(Prompt("Enter your age: ") + InputColor())
	var age int
	fmt.Scanf("%d", &age)
	fmt.Print(Out("Age("), In(age), Out(") = "), Result(GetAlphabeticNumber(age)), "\n")
}

func GetAlphabeticNumber(value int) string {
	if value == 40 {
		return "Fourty"
	} else {
		return "Unknown"
	}
}
