//go:generate go run github.com/99designs/gqlgen generate -v

package graph

// Resolver allows us to leverage DI to pass down things such as DB
type Resolver struct{}
