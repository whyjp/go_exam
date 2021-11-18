package main

import (
	"fmt"

	"wz/topicmodify/xmlhandler"
)

func main() {
	c := xmlhandler.NewXmlhandler("./Konsumer.xml")
	topic_bf, _ := c.GetTopic()
	fmt.Println("topic is now ", topic_bf)

	c.SetTopic("totopic")
	defer c.WriteCfg()

	topic, _ := c.GetTopic()
	fmt.Println("topic is now ", topic)
	fmt.Println("ServiceName is now ", c.ServiceName)
}
