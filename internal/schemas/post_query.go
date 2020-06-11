package schemas

import (
	"errors"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"github.com/patrickmn/go-cache"
	"time"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 08/06/2020 04:56
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (pq *PostQuery) RegistryQueryFunc(PostFunc, PostListFunc HandlerFuncQuery) {
	pq.getPostFunc = PostFunc
	pq.getPostListFunc = PostListFunc
}

func (pq *PostQuery) RegistryBatchFunc(UserPostFunc dataloader.BatchFunc) {
	if UserPostFunc != nil {
		pq.userPostLoader = dataloader.NewBatchedLoader(UserPostFunc, dataloader.WithCache(&Cache{cache.New(1*time.Minute, 1*time.Minute)}), dataloader.WithBatchCapacity(50))
	}
}

func (pq *PostQuery) InitSchema() error {
	var user = pq.initUserSchema()

	var queryPost = pq.initPostQuery(user)

	var err error
	pq.schema, err = graphql.NewSchema(
		graphql.SchemaConfig{Query: queryPost},
	)

	return err
}

func (pq *PostQuery) initUserSchema() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Name:        "User ID",
					Type:        graphql.String,
					Description: "ID of user",
				},
				"username": &graphql.Field{
					Name:        "UserName",
					Type:        graphql.String,
					Description: "User name",
				},
				"full_name": &graphql.Field{
					Name:        "FullName",
					Type:        graphql.String,
					Description: "Full Name",
				},
			},
			//IsTypeOf:    nil,
			//Description: "",
		},
	)

}

func (pq *PostQuery) initPostQuery(post *graphql.Object) *graphql.Object {
	var argsForPost = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:         graphql.String,
			DefaultValue: graphql.String,
			Description:  "An id",
		},
	}

	var argsForPosts = graphql.FieldConfigArgument{
		"ids": &graphql.ArgumentConfig{
			Type:         graphql.String,
			DefaultValue: graphql.String,
			Description:  "List id",
		},
	}

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:       "query",
			Interfaces: nil,
			Fields: graphql.Fields{
				"post": &graphql.Field{
					Name: "A Post",
					Type: post,
					Args: argsForPost,
					Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
						if pq.getPostFunc != nil {
							return pq.getPostFunc(p.Args)
						} else {
							return make(map[string]interface{}), errors.New(`Processing not registered`)
						}
					},
					DeprecationReason: "",
					Description:       "",
				},
				"posts": &graphql.Field{
					Name: "List Post",
					Type: post,
					Args: argsForPosts,
					Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
						if pq.getPostListFunc != nil {
							return pq.getPostListFunc(p.Args)
						} else {
							return make(map[string]interface{}), errors.New(`Processing not registered`)
						}
					},
					DeprecationReason: "",
					Description:       "",
				},
			},
			IsTypeOf:    nil,
			Description: "",
		},
	)
}

func (pq *PostQuery) GetSchema() graphql.Schema {
	return pq.schema
}
