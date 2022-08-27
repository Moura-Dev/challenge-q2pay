package models

type User struct {
	Id        int    `json:"-" db:"id"`
	FullName  string `json:"full_name" db:"full_name"`
	CpfCnpj   string `json:"cpfcnpj" db:"cpfcnpj"`
	Email     string `json:"email" db:"email"`
	Login     string `json:"login" db:"login"`
	Password  string `json:"password,omitempty" db:"password"`
	Active    bool   `json:"active" db:"active"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}
