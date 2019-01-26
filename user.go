package main

import (
	"errors"
	"time"
)

func (u User) register() (err error) {
	if u.Email == "" || u.Password == "" {
		return errors.New(" ID or Password is blank.")
	}

	query := `INSERT INTO users
(email, password, real_name, student_id, username, created_at, modified_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id`
	stmt, err := s.DB.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return
	}

	err = stmt.QueryRow(
		u.Email, u.Password, u.RealName, u.StudentID, u.UserName, time.Now(), time.Now()).Scan(u.ID)
	return
}

func login(email, password string) (success bool, err error) {
	query := "SELECT password FROM users WHERE email = $1"
	stmt, err := s.DB.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return
	}

	var encryptedPassword string
	err = stmt.QueryRow(email).Scan(encryptedPassword)
	if err != nil {
		return
	}

	success = comparePassword(encryptedPassword, password)

	return
}

func userByEmail(email string) (u User, err error) {
	query := `SELECT 
       id, password, real_name, student_id, username, created_at, modified_at, is_admin, is_certified, is_blocked 
FROM users WHERE email = $1`
	stmt, err := s.DB.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return
	}

	err = stmt.QueryRow(email).Scan(
		u.ID, u.Password, u.RealName, u.StudentID, u.UserName,
		u.CreatedAt, u.ModifiedAt, u.IsAdmin, u.IsCertified, u.IsBlocked)
	return
}

func userByID(id int) (u User, err error) {
	query := `SELECT email, password, real_name, student_id, username, created_at, modified_at FROM users
WHERE id = $1`
	stmt, err := s.DB.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return
	}

	err = stmt.QueryRow(id).Scan(
		u.Email, u.Password, u.RealName, u.StudentID, u.UserName, u.CreatedAt, u.ModifiedAt)
	return
}

func getUsers(limit, offset int) (users []User, err error) {
	query := `SELECT id, email, password, real_name, student_id, username, created_at, modified_at
FROM users ORDER BY id LIMIT $1 OFFSET $2 `

	stmt, err := s.DB.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return
	}

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return
	}

	for rows.Next() {

		err = rows.Scan() // TODO: 유저 풀 만들고 스캔하기
		if err != nil {
			return
		}

	}

	return
}

/*
func searchUser(key string, value interface{}) (users []User, err error) {

	format := `SELECT email, password, real_name, student_id, username, created_at, modified_at
FROM users WHERE %s = $1`

	switch key {
	case "id":
		i, ok := value.(int)
		if !ok {
			err = errors.New("ID should be integer")
			return
		}

		u, err := userByID(i)
		if err != nil {
			return
		}

		users = append(users, u)
		return

	case "created_at", "modified_at":
		i, ok := value.(time.Time)
		if !ok {
			err = errors.New("value should be time.Time")
			return
		}

		query := fmt.Sprintf(format, key)
		stmt, err := s.db.Prepare(query)
		defer closeStmt(stmt)
		if err != nil {
			return
		}

	}
	return
}
*/
