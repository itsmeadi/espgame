package main

import (
	"github.com/ESPGame/src/interfaces/api"
	"github.com/ESPGame/src/interfaces/db/mysql"
	"github.com/ESPGame/src/usecase"
	"log"
	"net/http"
)

func main() {

	repo := mysql.Conn

	uc := &usecase.QuestionsUseCase{
		UserRepo: repo,
		QuesRepo: repo,
		AnsRepo:  repo,
		QARepo:   repo,
	}

	interactor := api.Interactor{UseCase: uc}

	api := api.New(&api.API{Interactor: &interactor})
	api.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
