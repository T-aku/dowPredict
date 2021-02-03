package main

import (
	//"fmt"

	"./price"
	"./point"
	"./line"
)

func main() {
	dowClosePrice := price.GetDowCurrentPrice()
	price.InsertPriceToDB(dowClosePrice)
	dowOpenPrice := price.GetDowOpenPrice()

	votePredictOfTaguchi := point.GetVotePredictOfTaguchi()
	resultOfTaguchi := point.GetResult(dowClosePrice, dowOpenPrice, votePredictOfTaguchi)

	votePredictOfShimogai := point.GetVotePredictOfShimogai()
	resultOfShimogai := point.GetResult(dowClosePrice, dowOpenPrice, votePredictOfShimogai)
	
	point.InsertPointsToDb(resultOfTaguchi, resultOfShimogai)

	totalPointOfTaguchi := point.GetTotalPointOfTaguchi()
	totalPointOfShimogai := point.GetTotalPointOfShimogai()
	totalCounts := point.GetTotalCounts()

	line.SendResultMessage(dowOpenPrice, dowClosePrice, resultOfTaguchi, resultOfShimogai, totalPointOfTaguchi, totalPointOfShimogai, totalCounts)

	point.InsertVotesToDb()
 }