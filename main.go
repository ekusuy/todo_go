package main

import (
	"fmt"
	"go_todo/config"
)

func main() {
	fmt.Println(config.Config.Port)
}
