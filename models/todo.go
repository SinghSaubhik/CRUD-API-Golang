package models

type Todo struct {
	ID        string `bson:"_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
