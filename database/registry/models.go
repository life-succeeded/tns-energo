package registry

import "time"

type Item struct {
	Id            string    `bson:"_id,omitempty" json:"_id,omitempty"`
	AccountNumber string    `bson:"account_number" json:"account_number"`
	Surname       string    `bson:"surname" json:"surname"`
	Name          string    `bson:"name" json:"name"`
	Patronymic    string    `bson:"patronymic" json:"patronymic"`
	Object        string    `bson:"object" json:"object"`
	HaveAutomaton bool      `bson:"have_automaton" json:"have_automaton"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}
