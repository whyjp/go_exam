package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func StructPrintToJson(from interface{}) {
	data, _ := json.Marshal(from)

	fmt.Printf("%s \n", data)
}

func StringsToArray(from string) []string {
	arr := strings.Split(from, ";")
	return arr
}
