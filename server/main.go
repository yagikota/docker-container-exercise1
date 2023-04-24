// TCPサーバ server.go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	// https://pkg.go.dev/net#Listen
	listen, err := net.Listen("tcp", ":8080")
	add := listen.Addr().String()
	if err != nil {
		log.Fatalf("%sのリッスンに失敗しました\n", add)
	}
	defer listen.Close()
	fmt.Printf("%sで受付開始しました\n", add)

	// listen.Acceptは1回受け付けるとcloseしてしまうため、何度もAcceptを呼ぶ
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("コネクションを確立できませんでした")
		}

		buf := make([]byte, 1024)

		go func(conn net.Conn) {
			// リクエスト元のアドレス
			fmt.Printf("[Remote Address]\n%s\n", conn.RemoteAddr())

			// Readerを作成して、送られてきたメッセージを出力する
			n, _ := conn.Read(buf)
			fmt.Printf("[Message]\n%s", string(buf[:n]))

			time.Sleep(1 * time.Second)

			res := fmt.Sprintf("Hello, %s\n", conn.RemoteAddr())

			// レスポンスを書き込む
			// HTTP/1.1かつ、ContentLengthの設定が必要
			response := http.Response{
				StatusCode:    200,
				ProtoMajor:    1,
				ProtoMinor:    1,
				ContentLength: int64(len(res)),
				Body: io.NopCloser(
					strings.NewReader(res)),
			}
			response.Write(conn)
			// コネクションを切断する
			conn.Close()
		}(conn)
	}
}
