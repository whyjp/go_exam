package main

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

type Configuration struct {
	XMLName     xml.Name `xml:"configuration"`
	Comment     string   `xml:",comment"`
	Application struct {
		XMLName xml.Name `xml:"application"`
		Param   []Param  `xml:"param"`
	} `xml:"application"`
	Repository struct {
		XMLName xml.Name `xml:"repository"`
		Param   []Param  `xml:"param"`
	} `xml:"repository"`
	Log struct {
		XMLName  xml.Name   `xml:"log"`
		Appender []Appender `xml:"appender"`
	} `xml:"log"`
}

type Appender struct {
	XMLName xml.Name `xml:"appender"`
	Name    string   `xml:"name,attr"`
	Comment string   `xml:",comment"`
	Param   []Param  `xml:"param"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

func main() {
	// xml 파일 오픈
	fp, err := os.Open("e:/Konsumer.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, _ := ioutil.ReadAll(fp)

	// xml 디코딩
	var members Configuration
	xmlerr := xml.Unmarshal(data, &members)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Println(members)
	for i := range members.Repository.Param {
		if members.Repository.Param[i].Name == "topic" {
			fmt.Println("find topic : ", members.Repository.Param[i].Value)
		}
	}
	outdata, err_mar := xml.MarshalIndent(&members, "", "  ")
	if err_mar != nil {
		fmt.Println(err_mar)
		panic(outdata)
	}

	ioutil.WriteFile("e:/Konsumer_new.xml", outdata, fs.FileMode(os.O_APPEND))

	svc.
}
