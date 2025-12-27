// Package store will handle all relational database interactions so this
// is decoupled from the main business logic of the application
package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, post *Post) error
	}
	Users interface {
		Create(context.Context, user *User) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
