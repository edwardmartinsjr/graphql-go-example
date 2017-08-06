package types

import (
	"strconv"

	"github.com/graphql-go-example/model"
	"github.com/graphql-go-example/modules/fields"
	"github.com/graphql-go-example/modules/resolvers"

	"github.com/graphql-go/graphql"
)

//CommentType -
var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Comment",
	Fields: fields.CommentFields,
})

func init() {
	CommentType.AddFieldConfig("user", &graphql.Field{
		Type: UserType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if comment, ok := p.Source.(*model.Comment); ok == true {
				return resolvers.GetUserByID(comment.UserID)
			}
			return nil, nil
		},
	})
	CommentType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			i := p.Args["id"].(string)
			id, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				return nil, err
			}
			return resolvers.GetPostByID(id)
		},
	})
}
