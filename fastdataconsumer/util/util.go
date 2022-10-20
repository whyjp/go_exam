package util

import (
	"encoding/json"
	"log"
)

func LogingStructure(s interface{}) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Println("error:", err)
	}
	log.Print(string(b))
}
