// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

type Player struct {
	ID        int64
	UserID    int64
	Name      string
	BestScore int64
}

type User struct {
	ID           int64
	Username     string
	PasswordHash string
}
