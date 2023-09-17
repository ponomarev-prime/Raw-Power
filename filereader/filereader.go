package filereader

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFileEntries(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}

	lines := strings.Split(string(data), "\n")

	var names []string

	for _, line := range lines {
		// Удаление лишних пробелов и переносов строки.
		name := strings.TrimSpace(line)
		if name != "" {
			names = append(names, name, "|")
		}
	}

	return names
}

func arrayElementJoiner(stringArray []string) string {
	justString := strings.Join(stringArray, " ")
	return justString
}
