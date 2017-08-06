package fields

import (
	"github.com/graphql-go-example/model"
	"github.com/graphql-go/graphql"
)

//CommentFields -
var CommentFields = graphql.Fields{
	"id": &graphql.Field{
		Type: graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if comment, ok := p.Source.(*model.Comment); ok == true {
				return comment.ID, nil
			}
			return nil, nil
		},
	},
	"title": &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if comment, ok := p.Source.(*model.Comment); ok == true {
				return comment.Title, nil
			}
			return nil, nil
		},
	},
	"body": &graphql.Field{
		Type: graphql.NewNonNull(graphql.ID),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if comment, ok := p.Source.(*model.Comment); ok == true {
				return comment.Body, nil
			}
			return nil, nil
		},
	},
}
