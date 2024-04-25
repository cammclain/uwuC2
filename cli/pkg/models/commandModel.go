package models

type Command struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
	Status  string `json:"status"`
}
