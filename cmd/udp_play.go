package main

import (
	"fmt"

	"play-ground/module"
)

func main() {
	e := module.Err{Err: fmt.Errorf("error")}
	fmt.Println(e.Error())
}
