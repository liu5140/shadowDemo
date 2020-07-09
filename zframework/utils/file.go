package utils

import (
	"io/ioutil"
	"os"
)

func FilePutContents(path string, content string) {
	handle, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer handle.Close()
	if err == nil {
		handle.Write([]byte(content))
	}
}

func FileGetContents(path string) string {
	handle, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer handle.Close()
	buffer, err := ioutil.ReadAll(handle)
	if err != nil {
		panic(err)
	}
	return string(buffer)
}
