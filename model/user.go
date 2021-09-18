package model

type User struct {
	Model
	Nickname string `json:"nickname"`
}