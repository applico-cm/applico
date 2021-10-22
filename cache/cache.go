package cache

import "github.com/applico-cm/applico/model"

type Cache interface {
	GetUser(userID string) (*model.User, error)
	SetUser(user *model.User) error
}
