package main

import (
	"fmt"

	"github.com/Louis3-silv4/go-rest-api/database"
	"github.com/Louis3-silv4/go-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor ")
	routes.HandleRequest()
}
