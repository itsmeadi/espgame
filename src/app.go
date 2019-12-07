package main

import (
	api2 "github.com/ESPGame/src/interfaces/api"
	mysql2 "github.com/ESPGame/src/interfaces/db/mysql"
	usecase2 "github.com/ESPGame/src/usecase"
	"log"
	"net/http"
)

func main() {

	repo := mysql2.Conn

	uc := &usecase2.QuestionsUseCase{
		UserRepo: repo,
		QuesRepo: repo,
		AnsRepo:  repo,
		QARepo:   repo,
	}

	interactor := api2.Interactor{UseCase: uc}

	api := api2.New(&api2.API{Interactor: &interactor})
	api.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
