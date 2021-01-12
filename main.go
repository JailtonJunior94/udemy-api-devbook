package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JailtonJunior94/udemy-api-devbook/src/config"
	"github.com/JailtonJunior94/udemy-api-devbook/src/router"
)

func main() {
	config.Carregar()

	r := router.Gerar()

	fmt.Printf("🚀 API is running on http://localhost:%d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
