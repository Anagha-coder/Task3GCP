package models

// Employee struct with validations using the validator package
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      string `json:"role" validate:"required"`
	Deleted   bool   `json:"deleted"`
}
