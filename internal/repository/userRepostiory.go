package repository

import (
	"UserManagementProject/internal/domain"
	"database/sql"
)

type UserRepository interface {
	GetUser(id int) (domain.User, error)
	CreateUser(user domain.User) error
	UpdateUser(id domain.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user domain.User) error {
	statement, err := repo.db.Prepare("INSERT INTO users (ID,FirstName,LastName,Email ) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.ID, user.FirstName, user.LastName, user.Email)
	return err
}

func (repo *userRepository) GetUser(id int) (domain.User, error) {
	var user domain.User
	row := repo.db.QueryRow("SELECT ID, FirstName, LastName, Email FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return user, err
}

func (repo *userRepository) UpdateUser(user domain.User) error {
	statement, err := repo.db.Prepare("UPDATE users SET Firstname = ?, LastName = ?, Email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	return err
}

func (repo *userRepository) DeleteUser(id int) error {
	statement, err := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	return err
}

func (repo *userRepository) GetAllUsers() ([]domain.User, error) {
	rows, err := repo.db.Query("SELECT ID, FirstName,LastName,Email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
