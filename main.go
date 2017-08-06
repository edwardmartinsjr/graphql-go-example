package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go-example/conf"
	"github.com/graphql-go-example/modules"
	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
	_ "github.com/lib/pq"
)

func main() {

	//Create Schema
	var schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    modules.QueryType,
		Mutation: modules.MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}

	conf.LoadDBConfig()
	// simplest relay-compliant graphql server HTTP handler

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/graphql/v1", h)
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
