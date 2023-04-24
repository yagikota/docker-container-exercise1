package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", "server:8080")
	if err != nil {
		log.Fatal(err)
	}

	str := "Hello!"
	_, err = conn.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:count]))
}

func main() {
	http.HandleFunc("/send-request", handler)
	http.ListenAndServe(":8081", nil)
}
