package fields

import (
	"github.com/graphql-go-example/model"
	"github.com/graphql-go/graphql"
)

//UserFields -
var UserFields = graphql.Fields{
	"id": &graphql.Field{
		Type: graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return user.ID, nil
			}
			return nil, nil
		},
	},
	"email": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return user.Email, nil
			}
			return nil, nil
		},
	},
}
