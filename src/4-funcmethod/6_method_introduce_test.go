package funcmethod

import (
	"fmt"
	"math"
	"testing"
)

// 4.6.1 方法介绍
// 方法声明
type Point struct{ X, Y float64 }

// 对象作为接收器
// Distance same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
// 指针作为接收器
func (p *Point) DistancePtr(q *Point) float64 {
	if p == nil || q == nil {
		return 0
	}
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func Distance(q, p *Point)  {
	math.Hypot(q.X-p.X, q.Y-p.Y)
}
func TestDeclared(t *testing.T) {
	p := Point{
		X: 1,
		Y: 2,
	}
	q := &Point{
		X: 1,
		Y: 2,
	}
	p.DistancePtr(q)
	q.Distance(p)
}
func TestReceiverNil(t *testing.T) {
	var p *Point
	if p == nil {
		fmt.Println("p is nil")
	}
	p.DistancePtr(nil)
	// p.Distance(Point{}) // run err
}
