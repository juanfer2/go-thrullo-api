package repositories

import (
	"log"

	"github.com/juanfer2/go-thrullo-api.git/src/database"
)

func List(models []interface{}) interface{} {
	db := database.PopDb()
	err := db.All(&models)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(models)

	return models
}

type Builder interface {
	List(models []interface{}) Builder
}

// func (models interface{}) TypeInterface() {
// 	s := (models)(nil)
// 	t := reflect.TypeOf(s).Elem()
// }

type JoyDivision struct {
	Collection interface{}
}

func (b *JoyDivision) List(models []interface{}) Builder {
	return &JoyDivision{List(models)}
}
