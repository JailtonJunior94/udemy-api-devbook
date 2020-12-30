package router

import (
	"github.com/JailtonJunior94/udemy-api-devbook/src/router/rotas"
	"github.com/gorilla/mux"
)

// Gerar - retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
