package main

import (
	"golang2/controllers"
	"golang2/database"
	"golang2/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/index", controllers.HandlerIndex)
	router.HandleFunc("/login", controllers.HandlerLogin).Methods("POST")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "12345678",
			DB:         "db_golang2",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.User{})
}

// curl -X POST --user zikri:zikri123 http://localhost:8080/login
// {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzg2NjQ3NjYsImlzcyI6Ik15IFNpbXBsZSBKV1QgQXBwIiwiVXNlcm5hbWUiOiJub3ZhbCIsIkVtYWlsIjoidGVycGFsbXVyYWhAZ21haWwuY29tIiwiR3JvdXAiOiJhZG1pbiJ9.tEpRwudxmt4v-71UlWPoOe3IA_MIWWJjf1CzCIZcgnk"}

// curl -X GET \
//  --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzg2NjQ3NjYsImlzcyI6Ik15IFNpbXBsZSBKV1QgQXBwIiwiVXNlcm5hbWUiOiJub3ZhbCIsIkVtYWlsIjoidGVycGFsbXVyYWhAZ21haWwuY29tIiwiR3JvdXAiOiJhZG1pbiJ9.tEpRwudxmt4v-71UlWPoOe3IA_MIWWJjf1CzCIZcgnk" \
//  http://localhost:8080/index
