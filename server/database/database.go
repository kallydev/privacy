package database

import (
	"context"
)

type Database interface {
	QueryByQQNumber(ctx context.Context, qqNumber int64) ([]Model, error)
	QueryByEmail(ctx context.Context, email string) ([]Model, error)
	QueryByIDNumber(ctx context.Context, idNumber string) ([]Model, error)
	QueryByPhoneNumber(ctx context.Context, phoneNumber int64) ([]Model, error)
}

type Model interface {
	GetName() (name string, valid bool)
	GetNickname() (nickname string, valid bool)
	GetPassword() (password string, valid bool)
	GetEmail() (email string, valid bool)
	GetQQNumber() (qqNumber int64, valid bool)
	GetIDNumber() (idNumber string, valid bool)
	GetPhoneNumber() (phoneNumber int64, valid bool)
	GetAddress() (address string, valid bool)
}
