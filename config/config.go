package config

import (
	"fmt"
	"os"
)

var (
	TestFilePath string
)

func init() {
	TestFilePath = os.Getenv(TestFilePath)
	if TestFilePath == "" {
		TestFilePath = fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "/src/github.com/OrthantBing/query-builder-go/examples")
		fmt.Println(TestFilePath)
	}
}
