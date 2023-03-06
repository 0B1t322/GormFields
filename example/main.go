package main

import (
	"fmt"

	"github.com/0B1t322/GormFields/example/gorm"
	"github.com/0B1t322/GormFields/example/gorm/models"
	"github.com/google/uuid"
)

func main() {
	db, err := gorm.ConnectToSqlite()
	if err != nil {
		panic(err)
	}

	id := uuid.New()

	db.
		Save(models.Model{
			ID:    id,
			Title: "Some",
		})

	var finded models.Model

	db.Model(models.Model{}).
		Where(models.ModelIdField.EQ().WithTable(), id).
		First(&finded)

	fmt.Printf("%+v", finded)
}
