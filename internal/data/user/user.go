package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	// ErrNotFound is used when a specific Product is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")
)

// RepositoryDb defines the repository for the book service.
type RepositoryDb struct {
	Db *sqlx.DB
}

// Repo is the interface for the maybe repository.
type Repo interface{}

// New returns a pointer to a book repo.
func New(db *sqlx.DB) *RepositoryDb {
	return &RepositoryDb{Db: db}
}
