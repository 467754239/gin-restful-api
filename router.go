package main

import (
	. "github.com/467754239/db-api/apis"
	"gopkg.in/gin-gonic/gin.v1"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	// Person
	router.POST("/person", AddPersonApi)
	router.GET("/persons", GetPersonsApi)
	router.GET("/person/:id", GetPersonApi)
	router.PUT("/person/:id", ModPersonApi)
	router.DELETE("/person/:id", DelPersonApi)

	return router
}
