package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sonderkevin/gql/app/repositories"
	"github.com/sonderkevin/gql/app/services"
	"github.com/sonderkevin/gql/graph"
	"github.com/sonderkevin/gql/graph/generated"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port, db := setup()
	srv := setupServer(db)
	runServer(srv, port)
}

func setup() (string, *gorm.DB) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, pass, dbName)

	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	fmt.Println("Connected to the database!")

	return port, db
}

func setupServer(db *gorm.DB) *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					CategoriaService: &services.CategoriaService{
						Repository: &repositories.CategoriaRepository{
							DB: db,
						},
					},
				},
			},
		),
	)
}

func runServer(srv *handler.Server, port string) {
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	fmt.Println("Server listening on port :" + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
