package line

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"strconv"
	"math"
	//"fmt"
)


func SendResultMessage(dowOpenPrice float64, dowClosePrice float64, resultOfTaguchi int, resultOfShimogai int, totalPointOfTaguchi int, totalPointOfShimogai int, totalCounts int) {
	accessToken := "RzSwANF7FVsL3k3ZZDQmGJdsmHUYhp2grJ0iFQz9I6J"

	differenceOfDow := dowClosePrice - dowOpenPrice
	absDifferenceOfDow := math.Abs(differenceOfDow)
	absDifferenceOfDowString := strconv.FormatFloat(absDifferenceOfDow, 'f', 2, 64)

	dowRatio := math.Abs(differenceOfDow / dowOpenPrice * 100)
	dowRatioString := strconv.FormatFloat(dowRatio, 'f', 2, 64)

	var msg string

	if (differenceOfDow >= 0 ) {
		if (resultOfTaguchi == 1 && resultOfShimogai == 1) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)の上げちんこだったよ！！！" + "田口もしもがいも当たりだ〜"
		} else if (resultOfTaguchi == 1 && resultOfShimogai == 0) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)の上げちんこだったよ！！！" + "田口は当たったけどしもがいは爆死"
		} else if (resultOfTaguchi == 0 && resultOfShimogai == 1) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)の上げちんこだったよ！！！" + "しもがいは当たったけど田口は爆死"
		} else if (resultOfTaguchi == 0 && resultOfShimogai == 0) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)の上げちんこだったよ！！！" + "田口もしもがいも爆死"
		}
	} else if ( differenceOfDow < 0){ 
		if (resultOfTaguchi == 1 && resultOfShimogai == 1) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)のダウンだったよ！！！" + "田口もしもがいも当たりだ〜"
		} else if (resultOfTaguchi == 1 && resultOfShimogai == 0) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)のダウンだったよ！！！" + "田口は当たったけどしもがいは爆死"
		} else if (resultOfTaguchi == 0 && resultOfShimogai == 1) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)のダウンだったよ！！！" + "しもがいは当たったけど田口は爆死"
		} else if (resultOfTaguchi == 0 && resultOfShimogai == 0) {
			msg = "昨日ダウは" + absDifferenceOfDowString + "ドル(" + dowRatioString + "%)のダウンだったよ！！！" + "田口もしもがいも爆死"
		}
	}

	fullMessage := msg + "\n" +
								 "田口: " + strconv.Itoa(totalPointOfTaguchi) + "\n" +
								 "しもがい: " + strconv.Itoa(totalPointOfShimogai) + "\n" +
								 "合計: " + strconv.Itoa(totalCounts) + "回\n\n" +
								 "今日の投票はhttp://18.180.86.100:8080/ まで!!"


	URL := "https://notify-api.line.me/api/notify"

	u, err := url.ParseRequestURI(URL)
	if err != nil {
			log.Fatal(err)
	}

	c := &http.Client{}

	form := url.Values{}
	form.Add("message", fullMessage)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
			log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = c.Do(req)
	if err != nil {
			log.Fatal(err)
	}
}