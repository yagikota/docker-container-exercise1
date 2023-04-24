// // TCPクライアント client.go
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"net/http"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		log.Fatal("tcp://127.0.0.1:8080に接続できませんでした")
// 	}
// 	defer conn.Close()

// 	// // メッセージを送信する
// 	request, err := http.NewRequest(
// 		"GET", "http://localhost:8080", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	request.Write(conn)

// 	// メッセージを受信する
// 	response, err := http.ReadResponse(
// 		bufio.NewReader(conn), request)
// 	if err != nil {
// 		fmt.Println("Retry")
// 	}
// 	// 結果を表示
// 	b, _ := io.ReadAll(response.Body)
// 	fmt.Println(string(b))
// }

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "server:8888")
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
