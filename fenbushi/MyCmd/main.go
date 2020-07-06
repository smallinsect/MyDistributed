package main

import (
	"flag"
	"fmt"
)

func main() {
	//第一个参数，为参数名称，第二个参数为默认值，第三个参数是说明
	username := flag.String("name", "", "Input your username")
	flag.Parse()
	fmt.Println("Hello, ", *username)
}
