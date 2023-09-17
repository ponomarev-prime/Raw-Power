package main

import (
	"fmt"
	"os"
	"strings"
)

type CmdConfigs struct {
	fileName string
	flagX    bool
}

// Значения по умолчанию
var defaultCmdConfigs = CmdConfigs{
	fileName: "default.txt",
	flagX:    false,
}

func main() {
	fmt.Println("begin")

	fmt.Println(getCmdConfigs(getSysArgs()))

	stringArray := []string{"Hello", "world", "!", "!!"}
	fmt.Println(arrayElementJoiner(stringArray))
}

func getSysArgs() []string {
	argsWithoutProg := os.Args[1:]
	return argsWithoutProg
}

func arrayElementJoiner(stringArray []string) string {
	justString := strings.Join(stringArray, " ")
	return justString
}

func getCmdConfigs(sysArgs []string) CmdConfigs {
	if len(sysArgs) == 0 {
		fmt.Println("Zero args")
		CmdConfigs := defaultCmdConfigs
		return CmdConfigs
	}

	partName := strings.Split(sysArgs[0], "=")[1]

	CmdConfigs := defaultCmdConfigs
	CmdConfigs.fileName = partName
	CmdConfigs.flagX = true
	return CmdConfigs
}
