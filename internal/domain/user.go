package domain

type User struct {
	Id        int    `json:"id" db:"id""`
	Email     string `json:"email" db:"email"`
	Role      string `json:"role" db:"role"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}
