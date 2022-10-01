package main

type Point struct {
	x, y int
}
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.x = 8
	w.Circle.Center.y = 9
	//初始化
	w = Wheel{
		Circle: Circle{
			Center: Point{x: 1, y: 1},
			Radius: 1,
		},
		Spokes: 20,
	}
}
