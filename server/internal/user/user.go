package user

type User struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
