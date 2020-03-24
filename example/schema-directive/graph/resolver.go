package graph

import (
	"github.com/99designs/gqlgen/example/schema-directive/graph/generated"
	"github.com/99designs/gqlgen/example/schema-directive/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
var User = map[string]*model.User{
	"1": {
		ID:   "User:1",
		Name: "User_1",
		Role: model.RoleAdmin,
	},
	"2": {
		ID:   "User:2",
		Name: "User_2",
		Role: model.RoleOwner,
	},
}

func NewRootResolver() generated.ResolverRoot {
	return &Resolver{
		users: User,
		todos: []*model.Todo{
			{
				ID:   "Todo:1",
				Text: "Buy a cat food",
				Done: true,
				User: User["2"],
			},
			{
				ID:   "Todo:2",
				Text: "Check cat water",
				Done: false,
				User: User["2"],
			},
		},
	}
}

type Resolver struct {
	users map[string]*model.User
	todos []*model.Todo
}
