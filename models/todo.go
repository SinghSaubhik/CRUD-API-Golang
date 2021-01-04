package models

type Todo struct {
	ID        string `json:"id" bson:"_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
