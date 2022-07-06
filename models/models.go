package models

type Event struct {
	ID    string `json:"id,omitempty" bson:"_id"`
	Title string `json:"title" bson:"Title"`
	Date  string `json:"date" bson:"Date"`
	End   string `json:"end" bson:"End"`
	Color string `json:"color" bson:"Color"`
}
