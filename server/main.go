// TCPサーバ server.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	// TCP://127.0.0.1:8000でリッスン（hostを省略する場合は`:8888`のように指定する）
	// https://pkg.go.dev/net#Listen
	listen, err := net.Listen("tcp", ":8888")
	add := listen.Addr().String()
	if err != nil {
		log.Fatalf("%sのリッスンに失敗しました\n", add)
	}
	defer listen.Close()
	fmt.Printf("%sで受付開始しました\n", add)

	// listen.Acceptは1回受け付けるとcloseしてしまうため、何度もAcceptを呼ぶ
	for {
		// コネクションを確立する
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

			// メッセージを返す
			res := fmt.Sprintf("Hello, %s\n", conn.RemoteAddr())

			// conn.Write([]byte(res))
			// レスポンスを書き込む
			// HTTP/1.1かつ、ContentLengthの設定が必要
			response := http.Response{
				StatusCode:    200,
				ProtoMajor:    1,
				ProtoMinor:    1,
				ContentLength: int64(len(res)),
				Body: ioutil.NopCloser(
					strings.NewReader(res)),
			}
			response.Write(conn)
			// コネクションを切断する
			conn.Close()
		}(conn)
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"strings"
// 	"time"
// )

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Server is running at localhost:8888")
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			panic(err)
// 		}
// 		go func() {
// 			fmt.Printf("Accept %v\n", conn.RemoteAddr())
// 			// Accept後のソケットで何度も応答を返すためにループ
// 			for {
// 				// タイムアウトを設定
// 				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
// 				// リクエストを読み込む
// 				request, err := http.ReadRequest(bufio.NewReader(conn))
// 				if err != nil {
// 					// タイムアウトもしくはソケットクローズ時は終了
// 					// それ以外はエラーにする
// 					neterr, ok := err.(net.Error) // ダウンキャスト
// 					if ok && neterr.Timeout() {
// 						fmt.Println("Timeout")
// 						break
// 					} else if err == io.EOF {
// 						break
// 					}
// 					panic(err)
// 				}
// 				// リクエストを表示
// 				dump, err := httputil.DumpRequest(request, true)
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Println(string(dump))
// 				content := "Hello World\n"
// 				// レスポンスを書き込む
// 				// HTTP/1.1かつ、ContentLengthの設定が必要
// 				response := http.Response{
// 					StatusCode:    200,
// 					ProtoMajor:    1,
// 					ProtoMinor:    1,
// 					ContentLength: int64(len(content)),
// 					Body: ioutil.NopCloser(
// 						strings.NewReader(content)),
// 				}
// 				response.Write(conn)
// 			}
// 			conn.Close()
// 		}()
// 	}
// }
