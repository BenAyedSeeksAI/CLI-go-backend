package main

import (
	"fmt"
	"io/ioutil"
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
	return server
}
func checkMigration() bool {
	path, _ := os.Getwd()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}
	dbfile := "person.db"
	for _, file := range files {
		if file.Name() == dbfile {
			return true
		}
	}
	return false
}
func main() {
	if !checkMigration() {
		LoadDatabase()
		fmt.Println("Database created ...")
	}
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
			// display the persons list
			fmt.Println("The list is as following :")
			for _, element := range GetAllPersons() {
				fmt.Printf("[Name]: %v [Age]: %v\n", element.Name, element.Age)
			}
		} else if commands[0] == "runserver" {
			// run the server
			server := serverRunner()
			server.Run(":8099")
		}
	}

}
