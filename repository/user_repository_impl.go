package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(nim, email, name, password, role) VALUES (?, ?, ?, ?, ?)"

	var nim interface{} = user.NIM
	if nim == "" {
		nim = nil
	}

	result, err := tx.ExecContext(ctx, SQL, nim, user.Email, user.Name, user.Password, user.Role)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "SELECT id, nim, email, name, password, role FROM users WHERE nim = ? OR email = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.NIM, user.Email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.NIM, &user.Email, &user.Name, &user.Password, &user.Role)
		if err != nil {
			panic(err)
		}
		return user
	} else {
		panic("email/nim or password is incorrect")
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, nim, email, name, password, role FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	var nimPtr *string
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &nimPtr, &user.Email, &user.Name, &user.Password, &user.Role)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, nim, email, name, password, role FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.NIM, &user.Email, &user.Name, &user.Password, &user.Role)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindByRole(ctx context.Context, tx *sql.Tx, role bool) (domain.User, error) {
	SQL := "SELECT id, nim, email, name, password, role FROM users WHERE role = ?"
	rows, err := tx.QueryContext(ctx, SQL, role)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.NIM, &user.Email, &user.Name, &user.Password, &user.Role)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, nim string, email string) (domain.User, error) {
	SQL := "SELECT id, nim, email, name, password, role FROM users WHERE nim = ? OR email = ?"
	rows, err := tx.QueryContext(ctx, SQL, nim, email)
	helper.PanicIfError(err)
	defer rows.Close()

	var userNIM *string

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &userNIM, &user.Email, &user.Name, &user.Password, &user.Role)
		helper.PanicIfError(err)
		if userNIM != nil {
			user.NIM = *userNIM
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindByUsernamePublic(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, email, name, password, role FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, err
	}
}
