package model

import "time"

type Customer struct {
	CusteomerID string // unique
	Name        string
	// TODO
}

type User struct {
	CustomerID   string    `db:"customer_id"`
	UserID       string    `db:"id"` // unique
	Name         string    `db:"name"`
	RoleID       string    `db:"role_id"`
	Email        string    `db:"email"`
	CreationDate time.Time `db:"create_date"`
	UpdateDate   time.Time `db:"update_date"`
	Teams        []struct {
		TeamID string
		Name   string
	}
	// TODO
}

type Team struct {
	CustomerID string
	TeamID     string
	Name       string
	Users      []struct {
		UserID string
		Name   string
	}
	// TODO
}

type Role struct {
	RoleID string
	Name   string
}

type Feature struct {
	FeatureID string
	Roles     []string
}

type Project struct {
	ProjectID    string
	CustomerID   string
	TeamID       string
	Name         string
	Description  string
	State        string // Active, Archived, Deleted
	CreationDate time.Time
	UpdateDate   time.Time
}

type Template struct {
	TemplateID   string
	ProjectID    string
	FileName     string
	Type         string
	Descriptor   string
	CreationDate time.Time
	UpdateDate   time.Time
}
