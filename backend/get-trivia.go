package main

import (
		"fmt"
		"log"
		"time"
		"net/http"
		"io/ioutil"
		"encoding/json"
		"github.com/gorilla/mux"
)

const (
	host = "0.0.0.0"
	port = 8080
)

var listenHost string

type dailyTrivia struct {
		Day int `json:"day"`
		Month int `json:"month"`
		TriviaMsg string `json:"trivia"`
}

func todayTrivia(w http.ResponseWriter, r *http.Request) {
		today := time.Now()
		msg := getDailyTrivia("http://numbersapi.com/%d/%d/date", today)
		// TODO: Create error handling when the request return 502


		// handle json
		trivia := dailyTrivia{
				Day:today.Day(),
				Month:int(today.Month()),
				TriviaMsg:msg,
		}
		log.Printf("INFO: Getting trivia '%s'", msg)

		// json_data, err := json.Marshal(trivia)
		json.NewEncoder(w).Encode(trivia)
}

func getDailyTrivia(url_format string, date time.Time) string {
		_, m, d := date.Date()
		url := fmt.Sprintf(url_format, m, d)

		resp, err := http.Get(url)
		defer resp.Body.Close()

		if err != nil {
				log.Println("Error on getting trivia")
		}

		// if status code not successful, return empty string
		if resp.StatusCode != 200 {
				return ""
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
				log.Println("Error on reading trivia body")
		}
		return string(body)
}

func init()  {
		listenHost = fmt.Sprintf("%s:%d", host, port)
}

func main() {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/today", todayTrivia)
		log.Printf("Starting API on %s", listenHost)

		log.Fatal(http.ListenAndServe(listenHost, router))
}
