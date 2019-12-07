package api

import (
	"context"
	"encoding/json"
	"errors"
	models2 "github.com/ESPGame/src/entities/models"
	templatego2 "github.com/ESPGame/src/templatego"
	uuid "github.com/satori/go.uuid"
	"html"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type handler func(w http.ResponseWriter, r *http.Request)

type Response struct {
	Res  interface{} `json:"response"`
	Base Base        `json:"base"`
}

type Base struct {
	Error error
}

func (api *API) Wrapper(hand func(w http.ResponseWriter, r *http.Request) (interface{}, error)) handler {

	return handler(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			res, err := hand(w, r)

			resp := Response{}
			if err != nil {
				resp.Base.Error = err
			}
			resp.Res = res
			j, _ := json.Marshal(resp)
			_, err = w.Write(j)
			if err != nil {
				log.Println("Error while marshal", err)
			}
		})

}

func (api *API) Auth(hand handler) handler {

	return handler(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			cookie, err := r.Cookie("user_login_esp")
			ctx := r.Context()
			if err != nil || cookie == nil || cookie.Value == "" {
				//w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "text/html")

				http.Redirect(w, r, r.Host+"/signin", http.StatusSeeOther)
				time.Sleep(time.Second*5)
				return
			}
			ctx = context.WithValue(ctx, "user_id", cookie.Value)
			r=r.WithContext(ctx)
			hand(w, r)
		})

}

func sanitize(s string) string {
	return html.EscapeString(s)
}
func (api *API) SignIn(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()
	id := sanitize(r.FormValue("username"))
	pass := sanitize(r.FormValue("password"))

	user, err := api.Interactor.UseCase.SignIn(ctx, id, pass)

	if err != nil {
		return nil, errors.New("invalid user")
	}
	cookie := &http.Cookie{Name: "user_login_esp", Value: user.Id, HttpOnly: false}
	http.SetCookie(w, cookie)
	return user, err

}
func (api *API) SignOut(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	cookie := &http.Cookie{Name: "user_login_esp", MaxAge: -1}
	http.SetCookie(w, cookie)
	return "SignOut SuccessFull", nil

}

func (api *API) SignUp(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := context.Background()

	var user models2.User

	user.Id = sanitize(r.FormValue("username"))
	user.Password = sanitize(r.FormValue("password"))
	user.UserType = sanitize(r.FormValue("type"))
	user.Name = sanitize(r.FormValue("name"))

	err := api.Interactor.UseCase.
		SignUp(ctx, user)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{Name: "user_login_esp", Value: user.Id, HttpOnly: false}
	http.SetCookie(w, cookie)
	return nil, err

}

func (api *API) ShowScore(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	w.Header().Set("Content-Type", "text/html")

	uid,_:=ctx.Value("user_id").(string)

	score, err := api.Interactor.UseCase.GetScore(ctx, uid)
	if err!=nil{
		log.Println(err)
	}


	if err := templatego2.TemplateMap["score"].Execute(w, score); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}

}

func (api *API) ShowQuestion(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	w.Header().Set("Content-Type", "text/html")


	ques, _ := api.Interactor.UseCase.GetQuestionsAnswers(ctx)

	qtemplate := struct {
		Ques []models2.QuestionAnswersResponse
	}{
		Ques: ques,
	}

	if err := templatego2.TemplateMap["questions"].Execute(w, qtemplate); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}

}

func (api *API) GetQuestionAnswers(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()

	res, err := api.Interactor.UseCase.GetQuestionsAnswers(ctx)

	return res, err

}

func (api *API) InsertQuestion(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	return api.InsertQuestionAndUploadFile(ctx, r)
}

// This function returns the filename(to save in database) of the saved file
// or an error if it occurs
func (api *API) InsertQuestionAndUploadFile(ctx context.Context, r *http.Request) ([]string, error) {

	var fileNames []string
	err := r.ParseMultipartForm(64 << 20) // max size used by FormFile
	if err != nil {
		return fileNames, err
	}
	fhsQues := r.MultipartForm.File["ques_files"]
	if len(fhsQues) < 1 {
		return fileNames, errors.New("Invalid file question")
	}

	imageId, err := api.FileUpload(ctx, fhsQues[0])
	qid, err := api.Interactor.UseCase.InsertQuestion(ctx, "Match the image", "./upload/"+imageId)
	if err != nil {
		log.Println(err)
	}

	fhsAns := r.MultipartForm.File["ans_files"]
	for _, fh := range fhsAns {

		imageId, err := api.FileUpload(ctx, fh)
		_, err = api.Interactor.UseCase.InsertAns(ctx, qid, "./upload/"+imageId)
		if err != nil {
			log.Println(err)
			continue
		}

		fileNames = append(fileNames, fh.Filename)
	}

	return fileNames, nil
}

func (api *API) FileUpload(ctx context.Context, fh *multipart.FileHeader) (string, error) {
	uuid, _ := uuid.NewV4()
	imageId := uuid.String()
	file, err := fh.Open()
	if file != nil {
		defer file.Close()
	}
	if err != nil || file == nil {
		return imageId, err
	}
	f, err := os.OpenFile("./upload/"+imageId, os.O_WRONLY|os.O_CREATE, 0666)
	if f != nil {
		defer f.Close()
	}
	if err != nil || f == nil {
		return imageId, err
	}

	// Copy the file to the destination path
	_, err = io.Copy(f, file)
	if err != nil {
		return imageId, err
	}
	return imageId, nil
}

func(api *API) SubmitAns(w http.ResponseWriter, r *http.Request) {

	ctx:=r.Context()
	uid,_:=ctx.Value("user_id").(string)
	if err := r.ParseForm(); err != nil {
		// handle error
	}

	var qid, aid int64


	for key, values := range r.PostForm {
		ques:=strings.Split(key,"ques-")
		if len(ques)>1{
			qid,_=strconv.ParseInt(ques[1], 10, 64)
		}
		if len(values)<1{
			continue
		}
		ans:=strings.Split(values[0],"radio-")

		if len(ans)>1{
			aid,_=strconv.ParseInt(ans[1], 10, 64)
		}
		err := api.Interactor.UseCase.SaveAns(ctx, qid, aid, uid)
		if err!=nil{
			log.Println("ERR=",err)
		}
	}
}


func (api *API) GetUserScore(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()

	uid, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errors.New("invalid user")
	}
	return api.Interactor.UseCase.GetScore(ctx, uid)
}

func (api *API) SaveAns(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()

	uid, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errors.New("invalid user")
	}

	qid, err := strconv.ParseInt(sanitize(r.FormValue("qid")), 10, 64)
	if err != nil {
		return nil, errors.New("invalid input")
	}
	aid, err := strconv.ParseInt(sanitize(r.FormValue("aid")), 10, 64)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	err = api.Interactor.UseCase.SaveAns(ctx, qid, aid, uid)

	return nil, err

}

func (api *API) GetScore(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()

	uid, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errors.New("invalid user")
	}
	return api.Interactor.UseCase.GetScore(ctx, uid)
}

func (api *API) TestServer(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "Server Locked and Loaded", nil
}



// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}
