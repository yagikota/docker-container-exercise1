package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	url := "http://server:8080"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var res Response
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	time.Sleep(5 * time.Second)
	fmt.Println(res.Message)
}
