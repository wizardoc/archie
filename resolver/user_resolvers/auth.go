package user_resolvers

import (
	"context"
)

func (u User) Username(ctx context.Context) string {
	return "hello"
}
