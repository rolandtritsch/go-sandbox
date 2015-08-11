package main

import (
	"fmt"
	"github.com/rolandtritsch/go-sandbox-lib"
)

func main() {
	fmt.Println(play.Reverse("Hello, dnalor!"))

	fmt.Println(play.RunControl("Roland", 50, 50))
	fmt.Println(play.RunControl2("Roland", 50, 10))

	fmt.Println(play.RunSwitch(0))
	fmt.Println(play.RunSwitch2())

	fmt.Println(play.RunFunction(1))
	fmt.Println(play.RunFunction2("Roland"))
}
