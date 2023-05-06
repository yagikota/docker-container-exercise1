// TCPクライアント client.go
package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// )

// func main() {
// 	// tcp://127.0.0.1:8888に接続する
// 	conn, err := net.Dial("tcp", "localhost:8080")
// 	if err != nil {
// 		log.Fatalf("%sに接続できませんでした", conn.RemoteAddr().String())
// 	}
// 	defer conn.Close()

// 	// メッセージを送信する
// 	msg := fmt.Sprintf("Hello, %s\n", conn.RemoteAddr())
// 	conn.Write([]byte(msg))

// 	// メッセージを受信する
// 	res := make([]byte, 1024)
// 	n, _ := conn.Read(res)
// 	fmt.Println(string(res[:n]))
// }
