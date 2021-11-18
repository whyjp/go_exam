package model

import "encoding/xml"

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
