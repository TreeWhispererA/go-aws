package main

import (
	"fmt"
	"log"
	"os"

	"blockparty.co/test/routes"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	fmt.Println("Port Number is :" + port)
	color.Cyan("🌍Sever is running on localhost:" + port)

	router := gin.Default()
	router.Use(gin.Logger())
	routes.Route(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
