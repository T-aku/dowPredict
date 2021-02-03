package price

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/PuerkitoBio/goquery"
)

const url = "https://www.bloomberg.co.jp/quote/INDU:IND"

func GetDowCurrentPrice() float64 {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
 
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	dowPriceString := doc.Find(".price").Text()

	dowPriceStringWithoutComma := strings.Replace(dowPriceString, ",", "", 1)

	dowPriceFloat, _ := strconv.ParseFloat(dowPriceStringWithoutComma, 64)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	return
	// }
	return dowPriceFloat
 
 }

 func InsertPriceToDB(price float64) {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					UPDATE dowprice
					SET close = $1
					WHERE date = current_date -1;
				 `

	close := price

	_, err = Db.Exec(sql, close)
	if err != nil {
		fmt.Println(err)
	}
 }

 func GetDowOpenPrice() float64 {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					SELECT open
					FROM dowprice
					WHERE date = current_date -1;
				 `

	var dowOpenPrice float64 
	err = Db.QueryRow(sql).Scan(&dowOpenPrice)
	if err != nil {
		fmt.Println(err)
	}

	return dowOpenPrice

 }