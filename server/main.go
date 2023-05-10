// TCPサーバ server.go
package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"time"
// )

// func main() {
// 	// TCP://127.0.0.1:8000でリッスン（hostを省略する場合は`:8888`のように指定する）
// 	listen, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal("tcp://server:8080のリッスンに失敗しました")
// 	}
// 	defer listen.Close()
// 	fmt.Printf("%sで受付を開始しました", listen.Addr().String())

// 	// listen.Acceptは1回受け付けるとcloseしてしまうため、何度もAcceptを呼ぶ
// 	for {
// 		// コネクションを確立する
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			log.Fatal("コネクションを確立できませんでした")
// 		}

// 		buf := make([]byte, 1024)

// 		go func() {
// 			// リクエスト元のアドレス
// 			fmt.Printf("[Remote Address]\n%s\n", conn.RemoteAddr())

// 			// Readerを作成して、送られてきたメッセージを出力する
// 			n, _ := conn.Read(buf)
// 			fmt.Printf("[Message]\n%s", string(buf[:n]))

// 			time.Sleep(1 * time.Second)

// 			// メッセージを返す
// 			res := fmt.Sprintf("Hello, %s\n", conn.RemoteAddr())
// 			conn.Write([]byte(res))

// 			// コネクションを切断する
// 			conn.Close()
// 		}()
// 	}
// }
