package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func StructPrintToJson(from interface{}) {
	data, _ := json.MarshalIndent(from, "", " ")

	fmt.Printf("%s \n", data)
}

func StringsToArray(from string) []string {
	arr := strings.Split(from, ";")
	return arr
}

type ctxSetfunc func(key string, value interface{})

func ToContext(source map[string]interface{}, fun ctxSetfunc) {
	for key, val := range source {
		fun(key, val)
	}
}
