package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

//Función para abrir o crear nuestra base de datos SQLite

func InitDB() error {
	
	db, err := gorm.Open(sqlite.Open("escuela.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	//Guardamos la instancia en DB 
	DB = db

	//Crea tablas en caso de que estas no existan
	err = DB.AutoMigrate(&Student{}, &Subject{}, &Grade{})
	if err != nil {
		return err
	}

	fmt.Println("La base de datos se ha inicializado correctamente: escuela.db")

	return nil
}
