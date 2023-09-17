package main

import (
	"GoSysArgs/cmdln"
	"GoSysArgs/filereader"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("begin")

	appCmdArgs := cmdln.GetCmdConfigs(getSysArgs())
	fmt.Println(appCmdArgs)
	fmt.Println(appCmdArgs.PathFileName)

	fileEnt := filereader.GetFileEntries(appCmdArgs.PathFileName)
	fmt.Println(fileEnt)
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
