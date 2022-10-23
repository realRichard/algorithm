package exercise

import (
	"fmt"
	"sync"
)

// https://github.com/lifei6671/interview-go/blob/master/question/q001.md

// 交替打印数据和字母
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

// 解题思路
// 使用两个 chan, 分别用于通知对方进行打印
// 在主 goroutine 开启打印
// 同时使用 waitGroup 进行等待打印完成

func AlternatePrintAlphaNum() {
	number, alpha := make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				alpha <- struct{}{}
			}

		}
	}()
	wg.Add(1)
	go func() {
		letter := 'A'
		for {
			select {
			case <-alpha:
				if letter >= 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(letter))
				letter++
				fmt.Print(string(letter))
				letter++
				number <- struct{}{}
			}
		}
	}()
	number <- struct{}{}
	wg.Wait()
}
