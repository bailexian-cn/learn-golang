package __construction

import (
	"fmt"
	"testing"
)

// there is no while statement in go
func TestWhile(t *testing.T) {

}

func TestSwitch(t *testing.T) {
	i := 1
	var res string
	switch i {
	case 0:
	case 1:
		res = fmt.Sprintf("%s1", res)
	}
	if res != "1" {
		t.Errorf("except 1 but is %s", res)
	}
}
