package psql

import (
	"awasomeProject/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (b *Users) Create(ctx context.Context, user domain.User) error {
	_, err := b.db.Exec("INSERT INTO users (name, age, sex) values ($1, $2, $3)",
		user.Name, user.Age, user.Sex)

	return err
}

func (b *Users) GetByID(ctx context.Context, id int64) (domain.User, error) {
	var user domain.User
	err := b.db.QueryRow("SELECT id, name, age, sex FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Sex)
	if err == sql.ErrNoRows {
		return user, domain.ErrUserNotFound
	}

	return user, err
}

func (b *Users) GetAll(ctx context.Context) ([]domain.User, error) {
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

func (b *Users) Delete(ctx context.Context, id int64) error {
	_, err := b.db.Exec("DELETE FROM users WHERE id=$1", id)

	return err
}

func (b *Users) Update(ctx context.Context, id int64, inp domain.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *inp.Name)
		argId++
	}

	if inp.Age != nil {
		setValues = append(setValues, fmt.Sprintf("aAge=$%d", argId))
		args = append(args, *inp.Age)
		argId++
	}

	if inp.Sex != nil {
		setValues = append(setValues, fmt.Sprintf("sex=$%d", argId))
		args = append(args, *inp.Sex)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := b.db.Exec(query, args...)
	return err
}
