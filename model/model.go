package model

type Customer struct {
	CusteomerID string // unique
	Name        string
	// TODO
}

type User struct {
	CustomerID string
	UserID     string // unique
	Name       string
	RoleID     string
	Teams      []struct {
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
