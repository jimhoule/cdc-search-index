package dtos

type CreateUserDto struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}