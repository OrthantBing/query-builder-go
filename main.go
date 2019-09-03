package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	sql "github.com/OrthantBing/query-builder-go/sql"
)

func main() {
	jsonFile, err := os.Open("examples/1.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var rf sql.RuleFilter

	err = json.Unmarshal(byteValue, &rf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rf.Transform())
}
