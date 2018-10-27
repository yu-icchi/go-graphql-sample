package fields

import (
	"errors"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/yu-icchi/go-graphql-sample/model"
)

var UserField = &graphql.Field{
	Type:        UserType,
	Description: "Get a user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, ok := params.Args["id"].(string)
		if !ok {
			return nil, errors.New("not found")
		}
		return &model.User{
			ID:    id,
			Name:  "test",
			Email: "test@sample.com",
			Age:   23,
			Num:   time.Now().UnixNano(),
		}, nil
	},
}

var ListUsersField = &graphql.Field{
	Type:        graphql.NewList(UserType),
	Description: "list of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return []*model.User{
			{
				ID:    "1",
				Name:  "AAA",
				Email: "aaa@sample.com",
				Age:   12,
				Num:   time.Now().UnixNano(),
			},
			{
				ID:    "2",
				Name:  "BBB",
				Email: "bbb@sample.com",
				Age:   26,
				Num:   time.Now().UnixNano(),
			},
			{
				ID:    "3",
				Name:  "CCC",
				Email: "ccc@sample.com",
				Age:   33,
				Num:   time.Now().UnixNano(),
			},
		}, nil
	},
}

var CreateUserField = &graphql.Field{
	Type:        UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"age": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, _ := params.Args["name"].(string)
		email, _ := params.Args["email"].(string)
		arg, _ := params.Args["arg"].(int)

		user := model.User{
			ID:    "",
			Name:  name,
			Email: email,
			Age:   arg,
		}
		return user, nil
	},
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"num": &graphql.Field{
			Type: Int64Type,
		},
	},
})

var Int64Type = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Int64",
	Description: "int64",
	Serialize: func(value interface{}) interface{} {
		return value.(int64)
	},
	ParseValue: func(value interface{}) interface{} {
		return value.(int64)
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.ParseInt(valueAST.Value, 10, 64); err == nil {
				return intValue
			}
		}
		return nil
	},
})
