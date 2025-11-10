package main

import (
	"fmt"
	"go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando servidor REST com Go")
	routes.HandleRequest()
}
