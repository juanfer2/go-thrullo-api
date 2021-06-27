package repositories

import (
	"log"

	"github.com/juanfer2/go-thrullo-api.git/src/database"
	"github.com/juanfer2/go-thrullo-api.git/src/models"
)

//func CreateUser(inputUser models.User) {
func CreateUser() {
	db := database.PopDb()
	log.Print(db)
	board := models.Board{Title: "Mr.", Description: "Luke"}
	log.Print(board)

	_, err := db.ValidateAndSave(&board)

	if err != nil {
		log.Fatal(err)
	}
}
