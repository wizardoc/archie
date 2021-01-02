package resolver

import (
	"archie/resolver/user_resolvers"
	"context"
	"fmt"
)

type UserLoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Resolver) Login(ctx context.Context, params UserLoginParams) (*user_resolvers.User, error) {
	return nil, fmt.Errorf("hello man")

	//return &user_resolvers.User{}
}
