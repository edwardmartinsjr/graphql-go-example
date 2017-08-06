package types

import (
	"strconv"

	"github.com/graphql-go-example/model"
	"github.com/graphql-go-example/modules/fields"
	"github.com/graphql-go-example/modules/resolvers"
	"github.com/graphql-go/graphql"
)

//PostType -
var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Post",
	Fields: fields.PostFields,
})

func init() {
	PostType.AddFieldConfig("user", &graphql.Field{
		Type: graphql.NewNonNull(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				return resolvers.GetUserByID(post.UserID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comment", &graphql.Field{
		Type: CommentType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					return nil, err
				}
				return resolvers.GetCommentByIDAndPost(id, post.ID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comments", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CommentType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*model.Post); ok == true {
				return resolvers.GetCommentsForPost(post.ID)
			}
			return []model.Comment{}, nil
		},
	})
}
