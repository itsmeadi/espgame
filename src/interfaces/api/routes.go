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
	http.HandleFunc("/insert_question", api.Auth(api.Wrapper(api.InsertQuestion)))
	http.HandleFunc("/getqa", api.Auth(api.Wrapper(api.GetQuestionAnswers)))
	http.HandleFunc("/saveans", api.Auth(api.Wrapper(api.SaveAns)))
	http.HandleFunc("/getscore", api.Auth(api.Wrapper(api.GetUserScore)))
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))

	http.HandleFunc("/show_questions", api.Auth(api.ShowQuestion))
	http.HandleFunc("/submit_ans", api.Auth(api.SubmitAns))
	http.HandleFunc("/score", api.Auth(api.ShowScore))
	//http.HandleFunc("/score", api.Auth(api.ShowScore))

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))

	//http.HandleFunc("/", api.api.Wrapper(api.api.TestServer))
}

//
//func main() {
//
//	repo := mysql.Conn
//	uc := &usecase.QuestionsUseCase{
//		UserRepo: repo,
//		QuesRepo: repo,
//		AnsRepo:  repo,
//		QARepo:   repo,
//	}
//
//	interactor := Interactor{UseCase: uc}
//	api := New(&API{Interactor: &interactor})
//	api.InitRoutes()
//	log.Fatal(http.ListenAndServe(":8080", nil))
//
//}
