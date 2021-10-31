package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(fileName string) ([]string, bool) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, true
	}
	lines := strings.Split(string(data), "\n")
	return lines, false
}

func WriteFile(fileName string, set map[string]bool) bool {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		f.Close()
		return true
	}

	for key := range set {
		// fmt.Println(key)
		fmt.Fprintln(f, key)
		if err != nil {
			fmt.Println(err.Error())
			f.Close()
			return true
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
