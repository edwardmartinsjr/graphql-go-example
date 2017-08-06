package types

import (
	"strconv"

	"github.com/graphql-go-example/model"
	"github.com/graphql-go-example/modules/fields"
	"github.com/graphql-go-example/modules/resolvers"

	"github.com/graphql-go/graphql"
)

//UserType -
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "User",
	Fields: fields.UserFields,
})

func init() {
	UserType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					return nil, err
				}
				return resolvers.GetPostByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("posts", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PostType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return resolvers.GetPostsForUser(user.ID)
			}
			return []model.Post{}, nil
		},
	})
	UserType.AddFieldConfig("follower", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Follower ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					return nil, err
				}
				return resolvers.GetFollowerByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("followers", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return resolvers.GetFollowersForUser(user.ID)
			}
			return []model.User{}, nil
		},
	})
	UserType.AddFieldConfig("followee", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Followee ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					return nil, err
				}
				return resolvers.GetFolloweeByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("followees", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*model.User); ok == true {
				return resolvers.GetFolloweesForUser(user.ID)
			}
			return []model.User{}, nil
		},
	})
}
