package square

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	newX := s.start.x + int(s.a)
	newY := s.start.y - int(s.a)
	p := Point{x: newX, y: newY}

	return p
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}

// func main() {
// 	p := Point{x: 2, y: 3}
// 	s := Square{start: p, a: 12}

// 	fmt.Println(s.End())
// 	fmt.Println(s.Area())
// 	fmt.Println(s.Perimeter())
// }
