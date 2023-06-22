package funcmethod

import (
	"fmt"
	"testing"
)

// 4.7.1 方法值
func TestMethodVal(t *testing.T) {
	p := Point{
		X: 1,
		Y: 2,
	}
	q := Point{
		X: 1,
		Y: 2,
	}
	dist := p.Distance
	fmt.Println(dist(q))
}

// 4.7.2 方法表达式
func TestMethodExpression(t *testing.T) {
	p := Point{
		X: 1,
		Y: 2,
	}
	q := Point{
		X: 1,
		Y: 2,
	}
	dist := Point.Distance
	fmt.Println(dist(p, q))
}
