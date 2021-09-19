# Customers, users, roles and teams

## Customers
Applico is a multitenant application and supports multiple customers working together in separated enviroments.

## Users
Every customers can define its own list of users. Users can belong to different teams. 
The user has a unique role (administrator, manager, dev-ops, dev).

## Teams
Teams are groups of users. Teams can act on different environments and manage the environment variables.

## Roles
- Dev-ops can create environments, variables and secret. Dev-ops can execute deployments, but can't create releases
- Dev can create projects, template, releases and deploymnents. Dev can set environment variables, but can't see secrets
- Manager can act as a Dev and a Dev-ops. Manager can create teams
- Administrator, can manage every entities associated to a customer. Administrator can create users
