package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/example/schema-directive/graph/model"
	"github.com/99designs/gqlgen/graphql"

	"github.com/99designs/gqlgen/example/schema-directive/graph"
	"github.com/99designs/gqlgen/example/schema-directive/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dir := generated.DirectiveRoot{}
	dir.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
		if u := getCurrentUser(ctx); u == nil || !u.HasRole(role) {
			return nil, fmt.Errorf("Access denied")
		}
		return next(ctx)
	}
	dir.User = func(ctx context.Context, obj interface{}, next graphql.Resolver, id string) (res interface{}, err error) {
		return next(context.WithValue(ctx, "userId", id))
	}
	cfg := generated.Config{
		Resolvers:  graph.NewRootResolver(),
		Directives: dir,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getCurrentUser(ctx context.Context) *model.User {
	if id := getUserId(ctx); id != "" {
		return graph.User[id]
	}
	return nil
}

func getUserId(ctx context.Context) string {
	if id, ok := ctx.Value("userId").(string); ok {
		return id
	}
	return ""
}
