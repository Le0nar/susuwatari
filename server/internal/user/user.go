package user

type User struct {
	Name     string
	Position Position
}

type Position struct {
	X int
	Y int
}
