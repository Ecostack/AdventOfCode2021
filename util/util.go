package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetFileContentsSplit(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err", err)
		return nil, err
	}
	fullFile := string(content)
	fileSplit := strings.Split(fullFile, "\n")
	return fileSplit, nil
}
