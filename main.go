package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/yu-icchi/go-graphql-sample/fields"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			e := &ErrorResponse{
				Message: err.Error(),
			}
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ret := execQuery(string(body), Schema)
		json.NewEncoder(w).Encode(ret)
	})
	http.ListenAndServe(":8080", nil)
}

func execQuery(query string, schema graphql.Schema) *graphql.Result {
	ret := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(ret.Errors) > 0 {
		fmt.Printf("%v\n", ret.Errors)
	}
	return ret
}

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user":      fields.UserField,
		"listUsers": fields.ListUsersField,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": fields.CreateUserField,
	},
})
