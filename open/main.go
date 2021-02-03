package main

import (
	//"fmt"

	"./price"
)

func main() {
	dowPrice := price.GetDowPrice()
	//fmt.Println(dowPrice)
	price.InsertPriceToDB(dowPrice)
 }