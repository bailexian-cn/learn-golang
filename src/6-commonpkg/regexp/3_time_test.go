package regexp

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeOut(t *testing.T) {
	ticker := time.NewTicker(1*time.Second)
	i := 0
	go func(){
		for {
			select {
			case <- ticker.C:
				i++
				fmt.Println(i)
			}
		}
	}()
	select {
	case <- time.After(3*time.Second):
		fmt.Println("time out")
	}
	ticker.Stop()
}
