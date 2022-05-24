package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Course string  `json:"course"`
	Salary float64 `json:"salary"`
}

var people = []person{
	{ID: "1", Name: "Peixoto", Course: "BCC", Salary: 100},
	{ID: "2", Name: "Mirror", Course: "BCC", Salary: 1000},
	{ID: "3", Name: "SAJ", Course: "Eng", Salary: 10000},
}

func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

func createPeople(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON((&newPerson)); err != nil {
		return
	}

	for _, a := range people {
		if a.ID == newPerson.ID {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "person already created"})
			return
		}
	}

	people = append(people, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func findPerson(c *gin.Context) {
	id := c.Param("id")

	for _, a := range people {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func main() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.GET("/people/:id", findPerson)
	router.POST("/people", createPeople)
	router.Run("localhost:3001")
}
