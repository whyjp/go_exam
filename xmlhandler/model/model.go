package model

type Configuration struct {
	Application struct {
		Text  string `xml:",chardata"`
		Param []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"param"`
	} `xml:"application"`
	Repository struct {
		Text  string `xml:",chardata"`
		Param []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"param"`
	} `xml:"repository"`
	Log struct {
		Text     string `xml:",chardata"`
		Appender []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Param []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"param"`
		} `xml:"appender"`
	} `xml:"log"`
}
