package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("begin")

	printArgs := getSysArgs()
	fmt.Println(printArgs)

	stringArray := []string{"Hello", "world", "!", "!!"}
	printArrayJoined := arrayElementJoiner(stringArray)
	fmt.Println(printArrayJoined)
}

func getSysArgs() []string {
	argsWithoutProg := os.Args[1:]
	return argsWithoutProg
}

func arrayElementJoiner(stringArray []string) string {
	justString := strings.Join(stringArray, " ")
	return justString
}
