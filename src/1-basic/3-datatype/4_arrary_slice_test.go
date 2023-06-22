package __datatype

import (
	"fmt"
	"testing"
)

type Node struct {
	Name string
	Ip   string
}

func TestChangePointArr(t *testing.T) {
	var pointerArr = []*Node{
		{
			Name: "node1",
			Ip:   "192.168.122.1",
		},
		{
			Name: "node2",
			Ip:   "192.168.122.2",
		},
	}
	newIp := "xx.xx.xx.xx"
	name := "node1"
	for _, n := range pointerArr {
		if n.Name == name {
			n.Ip = newIp
		}
	}
	fmt.Printf("pointerArr after changed is:\n")
	for _, n := range pointerArr {
		if n.Name == name && n.Ip != newIp {
			t.Fatalf("except: { \"name\": \"node1\", \"ip\": \"xx.xx.xx.xx\"}")
		}
	}
}

func TestChangeObjectArr(t *testing.T) {
	var objectArr = []Node{
		{
			Name: "node1",
			Ip:   "192.168.122.1",
		},
		{
			Name: "node2",
			Ip:   "192.168.122.2",
		},
	}
	newIp := "xx.xx.xx.xx"
	name := "node1"
	for _, n := range objectArr {
		if n.Name == name {
			n.Ip = newIp
		}
	}
	fmt.Printf("pointerArr after changed is:\n")
	for _, n := range objectArr {
		if n.Name == name && n.Ip == newIp {
			t.Fatalf("except: { \"name\": \"node1\", \"ip\": \"192.168.122.1\"}")
		}
	}
}

func TestArrIndex(t *testing.T) {
	arr := []string{"aa", "bb", "cc", "aa"}
	fmt.Printf("arr: %v\n", arr)
}
