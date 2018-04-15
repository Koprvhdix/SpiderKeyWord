package main

import (
	"fmt"
	"SpiderKeyWord/config"
)

func main() {
	conf := config.LoadConfig()
	fmt.Println(conf)
}
