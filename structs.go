package main

import (
	"time"
)

type User struct {
	ID       int
	Password string
	Email    string

	RealName  string // real_name
	StudentID string

	UserName string // username

	CreatedAt  time.Time
	ModifiedAt time.Time

	IsAdmin     bool
	IsCertified bool
	IsBlocked   bool
}

type Session struct {
	ID     int
	UUID   string
	UserID int
	Email  string

	CreatedAt time.Time
}

type Thread struct {
	Title      string
	CreatedAt  time.Time
	ModifiedAt time.Time

	ThreadID int
	UserID   int

	NumUser int
	NumRes  int

	CategoryIDs []int

	upVote    int
	downVote  int
	TotalVote int

	Invisible bool
}

type Res struct {
	ThreadID int
	Order    int
	UserID   int

	Content string
	Date    time.Time

	upVote    int
	downVote  int
	TotalVote int

	Invisible bool
}

type Category struct {
	ID        int
	Name      string
	NumThread int
}
