package main

import (
	"log"
	"fmt"
	"net/http"
	"text/template"
	"net/url"
	"strings"

	"database/sql"
	_ "github.com/lib/pq"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func taguchiBullHandler(w http.ResponseWriter, r *http.Request) {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					UPDATE votes
					SET taguchi = 1
					WHERE date = current_date;
				 `

	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}

	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}

	accessToken := "RzSwANF7FVsL3k3ZZDQmGJdsmHUYhp2grJ0iFQz9I6J"

	fullMessage := "田口の今日の予想はあげ！！あげちんこ！！"

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


func taguchiBearHandler(w http.ResponseWriter, r *http.Request) {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					UPDATE votes
					SET taguchi = 0
					WHERE date = current_date;
				 `

	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}

	accessToken := "RzSwANF7FVsL3k3ZZDQmGJdsmHUYhp2grJ0iFQz9I6J"

	fullMessage := "田口の今日の予想はさげ！！さげまんこ・・・"

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


func shimogaiBullHandler(w http.ResponseWriter, r *http.Request) {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					UPDATE votes
					SET shimogai = 1
					WHERE date = current_date;
				 `

	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}

	accessToken := "RzSwANF7FVsL3k3ZZDQmGJdsmHUYhp2grJ0iFQz9I6J"

	fullMessage := "しもがいの今日の予想はあげ！！あげちんこ！！"

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


func shimogaiBearHandler(w http.ResponseWriter, r *http.Request) {
	Db, err := sql.Open("postgres", "host=172.19.0.2 user=postgres password=postgres dbname=taguchi sslmode=disable")
	defer Db.Close()
	if err != nil {
			log.Fatal(err)
	}
	sql := `
					UPDATE votes
					SET shimogai = 0
					WHERE date = current_date;
				 `

	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	t, err := template.ParseFiles("./template/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}

	accessToken := "RzSwANF7FVsL3k3ZZDQmGJdsmHUYhp2grJ0iFQz9I6J"

	fullMessage := "しもがいの今日の予想はさげ！！さげまんさげまん！"

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


func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/taguchibull", taguchiBullHandler)
	http.HandleFunc("/taguchibear", taguchiBearHandler)
	http.HandleFunc("/shimogaibull", shimogaiBullHandler)
	http.HandleFunc("/shimogaibear", shimogaiBearHandler)
	http.ListenAndServe(":8080", nil)
}