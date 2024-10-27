package main

import (
	"fmt"
	"time"
)

func main() {
	ping()
	go ping()
	done := make(chan bool)
	
	// 啟動不同的 slowPing 協程
	go slowPing(5, done)
	go slowPing(2, done)
	go slowPing(3, done)
	go slowPing(4, done)

	// [Base way]所有取值後關閉頻道會報錯
	func() {
		for i := 0; i < 4; i++ {
			value := <-done
			fmt.Println(value)
		}
		close(done)
	}()
	
	// [Other way]使用 range 來讀取通道中的訊號
	// for key := range done {
	// 	fmt.Println(key) // 打印從 done 通道中接收到的訊號
	// }

	ping()
}

func ping() {
	fmt.Println("pong!")
}

func slowPing(waitingSeconds int, doneChan chan bool) {
	time.Sleep(time.Duration(waitingSeconds) * time.Second)
	fmt.Printf("Slow %d pong!\n", waitingSeconds)
	doneChan <- true
	// [Other way]設計最後一個調用提交後關閉頻道
	// if waitingSeconds == 5 {
	// 	close(doneChan)
	// }
}
