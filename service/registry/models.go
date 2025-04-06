package registry

type Item struct {
	Id            string `json:"id"`
	AccountNumber string `json:"account_number"`
	Surname       string `json:"surname"`
	Name          string `json:"name"`
	Patronymic    string `json:"patronymic"`
	Object        string `json:"object"`
	HaveAutomaton bool   `json:"have_automaton"`
}
