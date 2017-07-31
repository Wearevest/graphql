package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/wearevest/graphql"
	"github.com/wearevest/graphql/testutil"
)

func main() {

	//h := handler.New(&handler.Config{
	//	Schema: &testutil.StarWarsSchema,
	//	Pretty: true,
	//})

	//http.Handle("graphql/", h)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        testutil.StarWarsSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={hero{name}}'")
	http.ListenAndServe(":8080", nil)
}
