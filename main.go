package main

import (
	"fmt"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando servidor REST com Go")
	routes.HandleRequest()
}
