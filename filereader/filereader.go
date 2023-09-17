package filereader

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFileEntries(filePath string) []byte {
	entries, err := os.ReadFile(filePath)
	check(err)
	return entries
}
