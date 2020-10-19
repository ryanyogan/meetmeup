//go:generate go run github.com/99designs/gqlgen generate -v

package graph

import "yogan.dev/meetmeup/pkg/postgres"

// Resolver allows us to leverage DI to pass down things such as DB
type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UserRepo
}
