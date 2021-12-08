package app

type Person struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
