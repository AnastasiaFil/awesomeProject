package psql

import (
	"awasomeProject/internal/domain"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (b *Users) Create(user domain.User) error {
	_, err := b.db.Exec("INSERT INTO users (name, age, sex) values ($1, $2, $3)",
		user.Name, user.Age, user.Sex)

	return err
}

func (b *Users) GetByID(id int64) (domain.User, error) {
	var user domain.User
	err := b.db.QueryRow("SELECT id, name, age, sex FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Sex)
	if err == sql.ErrNoRows {
		return user, domain.ErrUserNotFound
	}

	return user, err
}

func (b *Users) GetAll() ([]domain.User, error) {
	rows, err := b.db.Query("SELECT id, name, age, sex FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Sex); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, rows.Err()
}

func (b *Users) Delete(id int64) error {
	_, err := b.db.Exec("DELETE FROM users WHERE id=$1", id)

	return err
}

func (b *Users) Update(id int64, user domain.User) error {
	var user1 domain.User
	err := b.db.QueryRow("SELECT id, name, age, sex FROM users WHERE id=$1", id).
		Scan(&user1.ID, &user1.Name, &user1.Age, &user1.Sex)
	if err != nil {
		return err
	}

	if user.Name != "" {
		user1.Name = user.Name
	}
	if user.Age != 0 {
		user1.Age = user.Age
	}
	if user.Sex != "" {
		user1.Sex = user.Sex
	}

	_, err = b.db.Exec("UPDATE users SET name=$1, age=$2, sex=$3 WHERE id=$4",
		user1.Name, user1.Age, user1.Sex, id)

	return err
}
