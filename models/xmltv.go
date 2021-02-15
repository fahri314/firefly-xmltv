package models

import "encoding/xml"

/*Tv xmltv output scheme */
type Tv struct {
	XMLName           xml.Name `xml:"tv"`
	Text              string   `xml:",chardata"`
	GeneratorInfoName string   `xml:"generator-info-name,attr"`
	GeneratorInfoURL  string   `xml:"generator-info-url,attr"`
	Channel           []struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id,attr"`
		DisplayName struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"display-name"`
		Icon struct {
			Text string `xml:",chardata"`
			Src  string `xml:"src,attr"`
		} `xml:"icon"`
		URL string `xml:"url"`
	} `xml:"channel"`
	Programme []struct {
		Text    string `xml:",chardata"`
		Start   string `xml:"start,attr"`
		Stop    string `xml:"stop,attr"`
		Channel string `xml:"channel,attr"`
		Title   struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"title"`
		Desc struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"desc"`
	} `xml:"programme"`
}

/*Channel xmltv channel */
type Channel struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id,attr"`
	DisplayName struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"display-name"`
	Icon struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"icon"`
	URL string `xml:"url"`
}

/*Programme xmltv programme */
type Programme struct {
	Text    string `xml:",chardata"`
	Start   string `xml:"start,attr"`
	Stop    string `xml:"stop,attr"`
	Channel string `xml:"channel,attr"`
	Title   struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"title"`
	Desc struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"desc"`
}
