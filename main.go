package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Iniciamos la base de datos
	err := InitDB()
	if err != nil {
		log.Fatal("Hubo un error al inicializar la base de datos:", err)
	}

	//Cargamos las rutas
	r := NewRouter()

	//Mensaje que sera mostrado en consola al inicializar la base de datos
	fmt.Println("El servidor se encuentra funcionando en la siguiente pagina http://localhost:3030")

	//Colocamos el número del servidor en el que queremos que se inicie
	http.ListenAndServe(":3030", r)
}
