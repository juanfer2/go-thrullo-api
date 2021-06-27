package servers

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/juanfer2/go-thrullo-api.git/src/config"
	"github.com/juanfer2/go-thrullo-api.git/src/graph"
	"github.com/juanfer2/go-thrullo-api.git/src/graph/generated"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func StartServer() {
	config.Conn()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	//log.Fatal(http.ListenAndServe(":"+port, nil))
}
