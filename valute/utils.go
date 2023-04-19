package valute

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/CyclopsV/rur-quotes-cbrf/pkg"
)

func (v Valute) Hystory(date time.Time) (Hystory, error) {
	to := date.Format("02/01/2006")
	from := date.Add(-24 * 90 * time.Hour).Format("02/01/2006")
	url := fmt.Sprintf("https://www.cbr.ru/scripts/XML_dynamic.asp?date_req1=%v&date_req2=%v&VAL_NM_RQ=%v", from, to, v.Id)
	hRaw, err := pkg.Request(url)
	h := Hystory{}
	if err != nil {
		return h, err
	}
	err = pkg.XMLUnmarshalUnUTF(hRaw, &h)
	return h, err
}

func (v *Valute) MinMaxAvg(date time.Time) error {
	hs, err := v.Hystory(date)
	if err != nil {
		return err
	}
	sumPrice := 0.0
	for i, vh := range hs.Records {
		vh.Value = strings.Replace(vh.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(vh.Value, 64)
		if err != nil {
			continue
		}
		date, err := time.Parse("02.01.2006", vh.Date)
		if err != nil {
			continue
		}
		sumPrice += value
		if i == 0 {
			v.Min.Value = value
			v.Min.Date = date
			v.Max.Value = value
			v.Max.Date = date
			continue
		}
		if v.Min.Value > value {
			v.Min.Value = value
			v.Min.Date = date
		}
		if v.Max.Value < value {
			v.Max.Value = value
			v.Max.Date = date
		}
		if i+1 == len(hs.Records) {
			v.Avg = sumPrice / float64(i+1)
		}
	}
	return nil
}

func (v Valute) String() string {
	return fmt.Sprintf("%v (avg: %v)\n\tmax: %v date: %v\n\tmin: %v date: %v\n", v.Name, v.Avg, v.Max.Value, v.Max.Date, v.Min.Value, v.Min.Date)
}
