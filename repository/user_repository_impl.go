package repository

import (
	"context"
	"database/sql"
	"fmt"
	"survey/helper"
	"survey/model/domain"
	"time"
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

func (repository *UserRepositoryImpl) InsertResetPassword(ctx context.Context, tx *sql.Tx, token domain.ResetPassword) (domain.ResetPassword, error) {
	SQL := "insert into reset_password(user_id, token, expired_at) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, token.UserId, token.Token, token.Expired_at)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	token.Id = int(id)
	return token, nil
}

func (repository *UserRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, token string) (domain.ResetPassword, error) {
	SQL := "SELECT id, user_id, token, expired_at FROM reset_password WHERE token = ?"
	rows, err := tx.QueryContext(ctx, SQL, token)
	if err != nil {
		return domain.ResetPassword{}, fmt.Errorf("error querying token: %v", err)
	}
	defer rows.Close()

	resetPassword := domain.ResetPassword{}
	if rows.Next() {
		var expiredAt []uint8 // Use []uint8 to store datetime from database
		err := rows.Scan(&resetPassword.Id, &resetPassword.UserId, &resetPassword.Token, &expiredAt)
		if err != nil {
			return domain.ResetPassword{}, fmt.Errorf("error scanning token row: %v", err)
		}

		// Convert []uint8 (datetime from database) to time.Time
		resetPassword.Expired_at, err = time.Parse("2006-01-02 15:04:05", string(expiredAt))
		if err != nil {
			return domain.ResetPassword{}, fmt.Errorf("error parsing expired_at: %v", err)
		}

		return resetPassword, nil
	}

	return domain.ResetPassword{}, fmt.Errorf("token not found: %s", token)
}

func (repository *UserRepositoryImpl) DeletedByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "delete from reset_password where user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	if err != nil {
		panic(err)
	}

	user, err := repository.FindById(ctx, tx, userId)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "UPDATE users SET password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Id)
	if err != nil {
		fmt.Println("UpdatePassword SQL error:", err)
		return user, err
	}

	return user, nil
}
