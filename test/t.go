package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/sonderkevin/gql/app/repositories"
// 	"github.com/sonderkevin/gql/app/services"
// 	"github.com/sonderkevin/gql/graph"
// 	"github.com/sonderkevin/gql/graph/generated"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// )

// const defaultPort = "8080"

// func main() {
// 	port, db := setup()
// 	srv := setupServer(db)
// 	runServer(srv, port)
// }

// func setup() (string, *gorm.DB) {
// 	// if err := godotenv.Load(); err != nil {
// 	// 	panic(err)
// 	// }

// 	// port := os.Getenv("PORT")
// 	// dsn := os.Getenv("DSN")

// 	port := "8080"
// 	dsn := "host=localhost user=postgres password=rkvnh032!! dbname=wankar port=5432 sslmode=disable"

// 	if port == "" {
// 		port = defaultPort
// 	}

// 	db, err := gorm.Open(
// 		postgres.Open(dsn),
// 		&gorm.Config{
// 			Logger: logger.Default.LogMode(logger.Info),
// 		},
// 	)
// 	if err != nil {
// 		panic("failed to connect database" + err.Error())
// 	}

// 	return port, db
// }

// func setupServer(db *gorm.DB) *handler.Server {
// 	return handler.NewDefaultServer(
// 		generated.NewExecutableSchema(
// 			generated.Config{
// 				Resolvers: &graph.Resolver{
// 					CategoriaService: &services.CategoriaService{
// 						Repository: &repositories.CategoriaRepository{
// 							DB: db,
// 						},
// 					},
// 				},
// 			},
// 		),
// 	)
// }

// func runServer(srv *handler.Server, port string) {
// 	// http.Handle("/", http.FileServer(http.Dir("./public")))
// 	// http.Handle("/playground", playground.Handler("GraphQL playground", "/playground"))
// 	http.Handle("/graphql", srv)

// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }
