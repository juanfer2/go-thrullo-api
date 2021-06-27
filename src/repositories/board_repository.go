package repositories

import (
	"log"

	"github.com/juanfer2/go-thrullo-api.git/src/database"
	"github.com/juanfer2/go-thrullo-api.git/src/models"
)

//func CreateUser(inputUser models.User) {
func CreateBoard(inputBoard models.Board) (models.Board, error) {
	db := database.PopDb()
	_, err := db.ValidateAndSave(&inputBoard)

	if err != nil {
		log.Fatal(err)
	}

	return inputBoard, err
}

func ListBoards() []*models.Board {
	db := database.PopDb()
	boards := []*models.Board{}
	err := db.All(&boards)

	if err != nil {
		log.Fatal(err)
	}

	return boards
}

func FindByBoardId(id int) models.Board {
	db := database.PopDb()
	board := models.Board{}
	err := db.Find(&board, id)

	if err != nil {
		log.Fatal(err)
	}

	return board
}

func LastBoard() models.Board {
	db := database.PopDb()
	board := models.Board{}
	err := db.Last(&board)

	if err != nil {
		log.Fatal(err)
	}

	return board
}

func WhereBoard(searchByField interface{}) models.Board {
	db := database.PopDb()
	board := models.Board{}
	err := db.Last(&board)

	if err != nil {
		log.Fatal(err)
	}

	return board
}
