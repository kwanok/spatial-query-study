package location

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Location struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Point Point  `json:"point"`
}
