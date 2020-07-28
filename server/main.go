package main

import(
	"cms/generated/gqlgen"
	"cms/generated/prisma-client"
	"cms/middleware"
	resolver "cms/resolver"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	var opt *prisma.Options
	client := prisma.New(opt)

	resolver := resolver.Resolver{
		DB: client,
	}

	h := handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{
		Resolvers:  &resolver,
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()

	r.Use(middleware.CheckHeader())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":7010").Error()
}
