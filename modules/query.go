package modules

import (
	"strconv"

	"github.com/graphql-go-example/modules/types"
	"github.com/graphql-go-example/modules/resolvers"

	"github.com/graphql-go/graphql"
)

//QueryType -
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					return nil, err
				}
				return resolvers.GetUserByID(id)
			},
		},
	},
})
