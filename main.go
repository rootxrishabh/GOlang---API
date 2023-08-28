package main

import (
	"log"
	"github.com/rootxrishabh/GOlang---API/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.CallBookStore(r)
	log.Fatal(r.Run("localhost:8080"))
}
