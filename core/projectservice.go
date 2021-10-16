package core

import "github.com/applico-cm/applico/model"

type ProjectService interface {
	FindProjectsByCustomerID(customerID string) ([]*model.Project, error)
	FindProjectsByUser(user *model.User) ([]*model.Project, error)
	FindTemplatesByProjectID(projectID string) ([]*model.Template, error)
}
