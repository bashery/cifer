package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args

	name := arg[1]
	to := arg[2]
	num, _ := strconv.Atoi(to)

	for i := 0; i < num; i++ {
		//name = name[1:]
		name = strings.TrimPrefix(name, "q")
		name = "q" + name

	}

	fmt.Println(name)
}
