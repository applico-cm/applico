package store

import (
	"github.com/applico-cm/applico/model"
	dbr "github.com/gocraft/dbr/v2"
)

const (
	userTable = "user_account"
)

type UserStore struct {
	conn *dbr.Session
}

func NewUserStore(db *dbr.Session) *UserStore {
	return &UserStore{
		conn: db,
	}
}

func (u *UserStore) FindUserByID(userID string) (user *model.User, err error) {
	if _, err = u.conn.Select("*").From(userTable).Where("id = ?", userID).Load(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserStore) FindUsersByCustomerID(customerID string) (users []*model.User, err error) {
	if _, err = u.conn.Select("*").From(userTable).Where("id = ?", customerID).Load(&users); err != nil {
		return nil, err
	}
	return users, nil
}
