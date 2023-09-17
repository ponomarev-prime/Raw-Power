package cmdln

import "strings"

type CmdConfigs struct {
	PathFileName string
	FlagX        bool
	Arg3         string
	Arg4         string
}

var defaultCmdConfigs = CmdConfigs{
	PathFileName: "default.txt",
	FlagX:        false,
	Arg3:         "default-arg3",
	Arg4:         "default-arg4",
}

func GetCmdConfigs(sysArgs []string) CmdConfigs {
	// Копируем значения из defaultCmdConfigs
	configs := defaultCmdConfigs

	// Создаем карту для хранения аргументов и их значений
	argMap := make(map[string]string)

	for i := 0; i < len(sysArgs); i++ {
		arg := sysArgs[i]
		//fmt.Println(arg)
		//fmt.Println(arg[:2])
		if arg[:2] == "--" {
			printArg := arg[2:]
			splitArg := strings.SplitAfter(printArg, "=")
			//fmt.Println(splitArg)
			argName := splitArg[0]
			argMap[argName] = ""
			if len(splitArg) > 1 {
				argValue := splitArg[1]
				//argMap[argName] = argValue
				//fmt.Println(argValue)
				argMap[argName] = argValue
			}
			//fmt.Println(argName)

		}
	}

	//fmt.Println(argMap)
	// Переопределяем поля в configs, если они есть в аргументах
	if value, exists := argMap["names="]; exists {
		configs.PathFileName = value
	}
	if _, exists := argMap["flagx"]; exists {
		configs.FlagX = true
	}
	if value, exists := argMap["arg3="]; exists {
		configs.Arg3 = value
	}
	if value, exists := argMap["arg4="]; exists {
		configs.Arg4 = value
	}

	return configs
}
