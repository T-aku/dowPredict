package point

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"

)

func GetVotePredictOfTaguchi() int {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					SELECT taguchi
					FROM votes
					WHERE date = current_date -1;
				 `

	var votePredictOfTaguchi int 
	err = Db.QueryRow(sql).Scan(&votePredictOfTaguchi)
	if err != nil {
		fmt.Println(err)
	}

	return votePredictOfTaguchi
}

func GetVotePredictOfShimogai() int {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					SELECT shimogai
					FROM votes
					WHERE date = current_date -1;
				 `

	var votePredictOfShimogai int 
	err = Db.QueryRow(sql).Scan(&votePredictOfShimogai)
	if err != nil {
		fmt.Println(err)
	}

	return votePredictOfShimogai
}


func GetResult(dowClosePrice float64, dowOpenPrice float64, votePredict int) int {

	var Result int

	if (dowClosePrice - dowOpenPrice >= 0 && votePredict == 1) {
		Result = 1
	} else if (dowClosePrice - dowOpenPrice < 0 && votePredict == 1) {
		Result = 0
	} else if (dowClosePrice - dowOpenPrice >= 0 && votePredict == 0){
		Result = 0
	} else if (dowClosePrice - dowOpenPrice < 0 && votePredict == 0){
		Result = 1
	}

	return Result
	
}


func InsertPointsToDb(resultOfTaguchi int, resultOfShimogai int) {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					INSERT INTO points (date, taguchi, shimogai, total) 
					VALUES (current_date-1, $1, $2, 1);
				 `

	taguchi := resultOfTaguchi
	shimogai := resultOfShimogai

	_, err = Db.Exec(sql, taguchi, shimogai)
	if err != nil {
		fmt.Println(err)
	}
}


func GetTotalPointOfTaguchi() int {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					SELECT SUM(taguchi)
					FROM points;
				 `

	var totalPointOfTaguchi int 
	err = Db.QueryRow(sql).Scan(&totalPointOfTaguchi)
	if err != nil {
		fmt.Println(err)
	}

	return totalPointOfTaguchi
}


func GetTotalPointOfShimogai() int {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					SELECT SUM(shimogai)
					FROM points;
				 `

	var totalPointOfShimogai int 
	err = Db.QueryRow(sql).Scan(&totalPointOfShimogai)
	if err != nil {
		fmt.Println(err)
	}

	return totalPointOfShimogai
}


func GetTotalCounts() int {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					SELECT SUM(total)
					FROM points;
				 `

	var totalCounts int 
	err = Db.QueryRow(sql).Scan(&totalCounts)
	if err != nil {
		fmt.Println(err)
	}

	return totalCounts
}

func InsertVotesToDb() {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					INSERT INTO votes (date, taguchi, shimogai) 
					VALUES (current_date, 1, 1);
				 `

	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}