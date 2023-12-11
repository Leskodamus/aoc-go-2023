package util

type Point struct {
    X, Y int
}

func (p *Point) Add (o *Point) Point {
    return Point {p.X + o.X, p.Y + o.Y}
}

