package models

type User struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
