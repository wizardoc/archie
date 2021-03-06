package resolver

import (
	"archie/resolver/organization_resolver"
	"archie/resolver/user_resolvers"
)

type Resolver struct {
	user_resolvers.UserResolver
	organization_resolver.OrganizationResolver
}
