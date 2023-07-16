package repositories

import (
	"database/sql"
	"tsarka/models"
)

// Database represents an interface for working with the database
type Database interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// UserRepository represents the repository for working with the database
type UserRepository struct {
	db Database
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser creates a new user in the database and returns the generated ID
func (ur *UserRepository) CreateUser(user *models.User) (int, error) {
	var id int
	err := ur.db.QueryRow("INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id", user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetUserByID retrieves user information by ID
func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	row := ur.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates user information
func (ur *UserRepository) UpdateUser(user *models.User) error {
	_, err := ur.db.Exec("UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3", user.FirstName, user.LastName, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user by ID
func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
