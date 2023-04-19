package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CyclopsV/rur-quotes-cbrf/storage"
)

func main() {
	date := time.Now()
	valutes := storage.New()

	for _, v := range valutes.Valutes {
		if err := v.MinMaxAvg(date); err != nil {
			log.Printf("Не удалось получить данные о \"%v\"", v.Name)
			continue
		}
		log.Println(v)
	}
	fmt.Scanf(" ")
}
