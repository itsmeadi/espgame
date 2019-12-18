package api

import (
	usecase2 "github.com/ESPGame/src/usecase"
	"net/http"
)

type Interactor struct {
	UseCase usecase2.UseCaseInterface
}

//API is the api struct
type API struct {
	//Cfg        *config.MainConfig
	//AuthModule *auth.Module
	Interactor *Interactor
}

//New is the api initializer
func New(this *API) *API {
	return &API{Interactor: this.Interactor}
}

func (api *API) InitRoutes() {
	http.HandleFunc("/signin", api.SignIn)
	http.HandleFunc("/signup", api.Wrapper(api.SignUp))
	http.HandleFunc("/signout", api.Wrapper(api.SignOut))
	http.HandleFunc("/insert_question", api.Auth(api.InsertQuestion))
	http.HandleFunc("/insert_question_random", api.Auth(api.InsertQuestionRandom))
	http.HandleFunc("/getqa", api.Auth(api.Wrapper(api.GetQuestionAnswers)))
	http.HandleFunc("/saveans", api.Auth(api.Wrapper(api.SaveAns)))
	http.HandleFunc("/getscore", api.Auth(api.Wrapper(api.GetUserScore)))
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))

	http.HandleFunc("/show_questions", api.Auth(api.ShowQuestion))
	http.HandleFunc("/submit_ans", api.Auth(api.SubmitAns))
	http.HandleFunc("/score", api.Auth(api.ShowScore))

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))

	//http.HandleFunc("/", api.api.Wrapper(api.api.TestServer))
}
