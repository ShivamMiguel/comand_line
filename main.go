package main

import (
	"comand_line/app"
	"log"
	"os"
)

func main() {
	println("camada principal")

	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
