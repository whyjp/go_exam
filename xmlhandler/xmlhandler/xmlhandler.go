package xmlhandler

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"wz/topicmodify/xmlhandler/model"
)

type Xmlhandler struct {
	filepath    string
	topicP      *string
	ServiceName string
	members     model.Configuration
}

func NewXmlhandler(filepath_ string) *Xmlhandler {
	if filepath_ == "" {
		filepath_ = "e:/Konsumer.xml"
	}
	c := &Xmlhandler{}
	c.filepath = filepath_
	c.loadXml()
	return c
}

func (c *Xmlhandler) loadXml() {
	// xml 파일 오픈
	fp, err := os.OpenFile(c.filepath, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, _ := ioutil.ReadAll(fp)

	xmlerr := xml.Unmarshal(data, &c.members)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Println(c.members)

	for i := range c.members.Repository.Param {
		if c.members.Repository.Param[i].Name == "topic" {
			c.topicP = &c.members.Repository.Param[i].Value
		}
	}
	if c.topicP != nil {
		fmt.Println("find topic : ", *c.topicP)
	}

	for i := range c.members.Application.Param {
		if c.members.Application.Param[i].Name == "service_name" {
			c.ServiceName = c.members.Application.Param[i].Value
		}
	}

	fmt.Println("find service_name : ", c.ServiceName)
}

func (c *Xmlhandler) WriteCfg() {
	outdata, err_mar := xml.MarshalIndent(&c.members, "", "  ")
	if err_mar != nil {
		fmt.Println(err_mar)
		panic(outdata)
	}

	fp, err := os.OpenFile(c.filepath, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	fp.Seek(0, 0)
	bufWriter := bufio.NewWriter(fp)
	bufWriter.Write(outdata)
	bufWriter.Flush()
}

func (c *Xmlhandler) SetTopic(topic_ string) {
	*c.topicP = topic_
}

func (c *Xmlhandler) GetTopic() (string, error) {
	if c.topicP == nil {
		return "", fmt.Errorf("i hasn't topic")
	}
	return *c.topicP, nil
}
