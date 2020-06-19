package main

import (
	"log"

	"github.com/leonelrr12/twittar/bd"
	"github.com/leonelrr12/twittar/handlers"
)

func main() {
	if bd.CheckConnect() == false {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}
