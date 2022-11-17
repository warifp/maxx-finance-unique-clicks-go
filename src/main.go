package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	__refLink := os.Getenv("REFERRAL_LINK")
	if __refLink == "" {
		print("Please set REFERRAL_LINK in .env file")
		os.Exit(0)
	}

	results := make(chan string)
	for i := 0; i < 100; i++ {
		go func(num int) {
			results <- threadRequest(num, __refLink)
		}(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
}

func threadRequest(num int, refLink string) string {
	timeDelay := rand.Intn(1)
	time.Sleep(time.Duration(timeDelay) * time.Millisecond)
	if timeDelay%2 == 0 {

		url := refLink
		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\"")
		req.Header.Add("Accept", "*/*")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
		req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
		req.Header.Add("Sec-Fetch-Site", "same-origin")
		req.Header.Add("Sec-Fetch-Mode", "cors")
		req.Header.Add("Sec-Fetch-Dest", "empty")
		req.Header.Add("Accept-Encoding", "deflate")
		req.Header.Add("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
		req.Header.Add("Cookie", "")
		req.Header.Add("Content-Type", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)
		// return fmt.Sprintf("Response %d: %s", num, body)

		if body == nil {
			return fmt.Sprintf("Response %d: %s", num, "nil")
		} else if len(body) == 0 {
			return fmt.Sprintf("Response %d: %s", num, "empty")
		} else {
			// return fmt.Sprintf("Response %d: %s", num, string(body))
			return fmt.Sprintf("Thread [%d] Connected to Address [%s]", num, url)
		}
	}
	return fmt.Sprintf("Response %d: %s", num, "Error")
}
