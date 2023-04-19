package valute

import (
	"encoding/xml"
	"time"
)

type value struct {
	Value float64
	Date  time.Time
}

type Valute struct {
	Id      string `xml:"ID,attr"`
	Nominal int    `xml:"Nominal"`
	Name    string `xml:"Name"`
	Avg     float64
	Max     value
	Min     value
}

type Record struct {
	Date  string `xml:"Date,attr"`
	Value string `xml:"Value"`
}

type Hystory struct {
	XMLName xml.Name `xml:"ValCurs"`
	Records []Record `xml:"Record"`
}
