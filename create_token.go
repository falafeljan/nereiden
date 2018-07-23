package main

import (
	"errors"
	redis "github.com/go-redis/redis"
	"github.com/graphql-go/graphql"
)

func createTokenCreationMutation(client *redis.Client) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
		Args: graphql.FieldConfigArgument{
			"containerID": &graphql.ArgumentConfig{
				Description: "The ID of the controlled container.",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"imageTag": &graphql.ArgumentConfig{
				Description: "The ID of the image the container is locked onto.",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			containerID, ok := params.Args["containerID"].(string)
			if !ok {
				return nil, errors.New("`containerID` is expected to be a string")
			}

			imageTag, ok := params.Args["imageTag"].(string)
			if !ok {
				return nil, errors.New("`imageTag` is expected to be a string")
			}

			conf := TokenConf{
				ContainerID: containerID,
				ImageTag:    imageTag,
			}

			return createToken(client, &conf)
		},
	}
}
