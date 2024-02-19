package message

type BasicMessageDto struct {
	Type string
}

type AddUserDto struct {
	Name string
	Type string
}

type ChangePositionDto struct {
	Name      string
	Direction string
	Type      string
}
