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
	http.HandleFunc("/signin", Wrapper(SignIn))
	http.HandleFunc("/signup", Wrapper(SignUp))
	http.HandleFunc("/signout", Wrapper(SignOut))
	http.HandleFunc("/insert_question", Auth(Wrapper(InsertQuestion)))
	http.HandleFunc("/getqa", Auth(Wrapper(GetQuestionAnswers)))
	http.HandleFunc("/saveans", Auth(Wrapper(SaveAns)))
	http.HandleFunc("/getscore", Auth(Wrapper(GetUserScore)))
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))



	http.HandleFunc("/show_questions", Auth(ShowQuestion))
	http.HandleFunc("/submit_ans", Auth(SubmitAns))
	http.HandleFunc("/score", Auth(ShowScore))






	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))

	//http.HandleFunc("/", api.Wrapper(api.TestServer))
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
