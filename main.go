package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // seu front local
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", config.Porta),
			cors(r), // aqui envolvemos o router no CORS
		),
	)
}
