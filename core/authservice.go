package core

import "github.com/applico-cm/applico/model"

type Any interface{}

type AuthorizationService interface {
	Login(userID, password string) (user *model.User, token string, err error)
	ValidateToken(token string) (userID string, err error)
	Authorize(user *model.User, feat *model.Feature, entities []Any) bool
}
