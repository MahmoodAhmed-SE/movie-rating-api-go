package models

type USER struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}