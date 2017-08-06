package fields

import (
	"github.com/graphql-go-example/model"
	"github.com/graphql-go/graphql"
)

//PostFields -
var PostFields = graphql.Fields{
	"id": &graphql.Field{
		Type: graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				return post.ID, nil
			}
			return nil, nil
		},
	},
	"title": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				return post.Title, nil
			}
			return nil, nil
		},
	},
	"body": &graphql.Field{
		Type: graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				return post.Body, nil
			}
			return nil, nil
		},
	},
}
