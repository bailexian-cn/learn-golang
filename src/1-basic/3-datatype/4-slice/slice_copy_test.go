package __slice

import (
	"fmt"
	"testing"
)

type Node struct {
	Name string
	Ip   string
}

var objectArr = []Node{
	{
		Name: "node1",
		Ip: "192.168.122.1",
	},
	{
		Name: "node2",
		Ip: "192.168.122.2",
	},
}

var pointerArr = []*Node {
	{
		Name: "node1",
		Ip: "192.168.122.1",
	},
	{
		Name: "node2",
		Ip: "192.168.122.2",
	},
}

func TestChangePointerArr(t *testing.T) {
	newIp := "xx.xx.xx.xx"
	name := "node1"
	for _, n := range pointerArr {
		if n.Name == name {
			n.Ip = newIp
		}
	}
	fmt.Printf("pointerArr after changed is:\n")
	for _, n := range pointerArr {
		fmt.Printf("%v\n", n)
	}
}
/*Output:
&{node1 xx.xx.xx.xx}
&{node2 192.168.122.2}
*/

func TestChangeObjectArr(t *testing.T) {
	newIp := "xx.xx.xx.xx"
	name := "node1"
	for _, n := range objectArr {
		if n.Name == name {
			n.Ip = newIp
		}
	}
	fmt.Printf("pointerArr after changed is:\n")
	for _, n := range pointerArr {
		fmt.Printf("%v\n", n)
	}
}
/*Output:
&{node1 192.168.122.1}
&{node2 192.168.122.2}
 */