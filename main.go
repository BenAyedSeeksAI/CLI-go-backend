package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func getPersons(context *gin.Context) {
// 	personData := getData()
// 	context.IndentedJSON(http.StatusOK, personData)
// }

func getPersons(context *gin.Context) {
	persons := GetAllPersons()
	context.IndentedJSON(http.StatusOK, persons)
}
func test(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "OK !!",
	})
}
func serverRunner() *gin.Engine {
	server := gin.Default()
	// get all the persons
	server.GET("/persons", getPersons)
	// testing the server
	server.GET("/test", test)
	// // server.GET("/data", getPersons)
	return server
}
func main() {
	if len(os.Args) > 1 {
		commands := os.Args[1:]
		if commands[0] == "add" {
			// server := serverRunner()
			// server.Run(":8099")
			// create a record
			Name, Age := commands[1], commands[2]
			newage, err := strconv.Atoi(Age)
			if err != nil {
				panic(err.Error())
			}
			MakeRecord(&Person{Name: Name, Age: newage})

		} else if commands[0] == "display" {
			server := serverRunner()
			server.Run(":8099")
			// display the persons list
		}
	}

}
