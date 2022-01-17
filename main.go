package main

import (
	"os"

	"github.com/brunoquindeler/go-cli/app"
	"github.com/sirupsen/logrus"
)

func main() {
	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
