package util

import (
	"encoding/json"
	"fmt"
)

func StructPrintToJson(from interface{}) {
	data, _ := json.Marshal(from)

	fmt.Printf("%s \n", data)
}
