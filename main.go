package main

import (
	"GoSysArgs/cmdln"
	"GoSysArgs/filereader"
	"fmt"
	"os"
)

func main() {
	fmt.Println("begin")

	appCmdArgs := cmdln.GetCmdConfigs(getSysArgs())
	fmt.Println(appCmdArgs)
	fmt.Println(appCmdArgs.PathFileName)

	fileEnt := filereader.GetFileEntries(appCmdArgs.PathFileName)
	fmt.Println(fileEnt)
}

func getSysArgs() []string {
	argsWithoutProg := os.Args[1:]
	return argsWithoutProg
}
