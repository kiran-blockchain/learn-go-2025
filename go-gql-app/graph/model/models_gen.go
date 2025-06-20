// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewProduct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       string  `json:"price"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"createdAt,omitempty"`
	DeletedAt   *string `json:"deletedAt,omitempty"`
	UpdatedAt   *string `json:"updatedAt,omitempty"`
}

type Query struct {
}

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	CreatedAt *string `json:"createdAt,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}
