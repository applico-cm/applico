package core

import "github.com/applico-cm/applico/model"

type CustomerService interface {
	FindCustomerByID(customerID string) (*model.Customer, error)
	ListRoles() ([]*model.Role, error)
}

type UserService interface {
	FindUserByID(userID string) (*model.User, error)
	FindUsersByCustomerID(customerID string) ([]*model.User, error)
}

type TeamService interface {
	FindTeamByID(teamID string) (*model.Team, error)
	FindTeamsByCustomerID(customerID string) ([]*model.Team, error)
}
