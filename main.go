package main

import (
	"fmt"
)

func main() {
	fmt.Println(Prompt("My Web Store in Golang"))
	fmt.Println(Out("Output: ", GreenText("Green"), BoldText("White"), RedBoldText("Red")))
}
