package main

import (
	"GoSysArgs/cmdln"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("begin")

	appCmdArgs := cmdln.GetCmdConfigs(getSysArgs())
	fmt.Println(appCmdArgs)
	fmt.Println(appCmdArgs.PathFileName)

	//stringArray := []string{"Hello", "world", "!", "!!"}
	//fmt.Println(arrayElementJoiner(stringArray))
}

func getSysArgs() []string {
	argsWithoutProg := os.Args[1:]
	return argsWithoutProg
}

func arrayElementJoiner(stringArray []string) string {
	justString := strings.Join(stringArray, " ")
	return justString
}
