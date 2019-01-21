package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func (u *User) newSession() (sess Session, err error) {
	query := `INSERT into sessions (uuid, email, user_id, created_at)
VALUES ($1, $2, $3, $4) RETURNING (id, uuid, email, user_id, created_at)`

	stmt, err := s.db.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return
	}

	err = stmt.QueryRow(createUUID(), u.Email, u.ID, time.Now()).Scan(
		sess.ID, sess.UUID, sess.Email, sess.UserID, sess.CreatedAt)
	return
}
