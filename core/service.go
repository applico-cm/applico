package core

import (
	"fmt"

	"github.com/applico-cm/applico/model"
	"github.com/applico-cm/applico/store"
)

type CoreService interface {
	UserService
}

type Service struct {
	userStore store.UserStore
}

func NewCoreService() CoreService {
	s := &Service{
		// init
	}
	return s
}

func (s *Service) FindUsersByCustomerID(customerID string) ([]*model.User, error) {
	// validate parameters

	return s.userStore.FindUsersByCustomerID(customerID)
}

func (s *Service) FindUserByID(userID string) (*model.User, error) {
	return nil, fmt.Errorf("not implemented")
}
