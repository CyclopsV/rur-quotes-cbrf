package storage

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CyclopsV/rur-quotes-cbrf/pkg"
	"github.com/CyclopsV/rur-quotes-cbrf/valute"
)

type Valutes struct {
	XMLName xml.Name         `xml:"ValCurs"`
	Valutes []*valute.Valute `xml:"Valute"`
}

func New() *Valutes {
	today := time.Now()
	url := fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?date_req=%v", today.Format("02/01/2006"))
	vsRaw, err := pkg.Request(url)
	if err != nil {
		log.Printf("Ошибка запроса списка валют: \n%v", err)
		os.Exit(1)
	}
	vs := Valutes{}
	if err := pkg.XMLUnmarshalUnUTF(vsRaw, &vs); err != nil {
		log.Printf("Ошибка парсинга списка валют: \n%v", err)
		os.Exit(1)
	}
	return &vs
}
