package resolver

import "context"

func (r *Resolver) Hello(ctx context.Context) string {
	return "hello"
}
