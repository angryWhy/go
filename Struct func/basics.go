package main

type Point struct {
	x, y int
}

func main() {

}
func sum(p, q Point) int {
	z := p.x + q.x
	return z
}

// 参数p成为方法的接收者
func (p Point) center(x int) int {
	return p.x + x
}
