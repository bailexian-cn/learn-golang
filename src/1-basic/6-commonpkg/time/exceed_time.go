package main

import (
	"fmt"
	"time"
)

func main() {
	workDoneCh := make(chan struct{}, 1)
	cancel := make(chan struct{}, 1)
	go func() {
		LongTimeWork(cancel, workDoneCh)  //这是我们要控制超时的函数
		workDoneCh <- struct{}{}
	}()

	select { //下面的case只执行最早到来的那一个
	case <- workDoneCh: //LongTimeWork运行结束
		fmt.Printf("%s LongTimeWork return\n", time.Now())
	case <- time.After(20 * time.Second):    //timeout到来
		cancel <- struct{}{}
		fmt.Printf("%s LongTimeWork timeout\n", time.Now())
	}

	time.Sleep(1 * time.Minute)
}

func LongTimeWork(cancel chan struct{}, workDoneCh chan struct{}) {
    for {
		select {
		case <- cancel:
			return
		default:
			fmt.Printf("%s Working\n", time.Now())
			time.Sleep(3 * time.Second)
		}
	}
}

func exceedTime(max time.Duration, f func()) error {
	return nil
}
