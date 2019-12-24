package api

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/ESPGame/src/entities/constant"
	"github.com/ESPGame/src/entities/models"
	templatego2 "github.com/ESPGame/src/templatego"
	"github.com/gofrs/uuid"
	"html"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type handler func(w http.ResponseWriter, r *http.Request)

type Response struct {
	Res  interface{} `json:"response"`
	Base Base        `json:"base"`
}

type Base struct {
	Error error
}

var ErrUnAuthorized = errors.New("unauthorized")

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

				http.Redirect(w, r, r.Host+"/login.html", http.StatusSeeOther)
				return
			}
			ctx = context.WithValue(ctx, "user_id", cookie.Value)
			r = r.WithContext(ctx)
			hand(w, r)
		})

}

func sanitize(s string) string {
	return html.EscapeString(s)
}
func (api *API) SignIn(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := sanitize(r.FormValue("username"))
	pass := encrypt(r.FormValue("password"))

	user, err := api.Interactor.UseCase.SignIn(ctx, id, pass)

	if err != nil {
		w.Write([]byte("invalid user"))
		return
		//return nil, errors.New("invalid user")
	}

	cookie := &http.Cookie{Name: "user_login_esp", Value: user.Id, HttpOnly: false}
	http.SetCookie(w, cookie)
	if user.UserType == constant.UserRoleAdmin {
		http.Redirect(w, r, "/insert_question", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/show_questions", http.StatusSeeOther)
	}
	return
}
func (api *API) SignOut(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	cookie := &http.Cookie{Name: "user_login_esp", MaxAge: -1}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login.html", http.StatusSeeOther)
	return "SignOut SuccessFull", nil

}

func encrypt(inp string) string {
	hash := md5.Sum([]byte(inp))
	return hex.EncodeToString(hash[:])
}

func (api *API) SignUp(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := context.Background()

	var user models.User

	user.Id = sanitize(r.FormValue("username"))
	user.Password = encrypt(r.FormValue("password"))
	user.UserType = sanitize(r.FormValue("type"))
	user.Name = sanitize(r.FormValue("name"))

	err := api.Interactor.UseCase.
		SignUp(ctx, user)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{Name: "user_login_esp", Value: user.Id, HttpOnly: false}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login.html", http.StatusSeeOther)
	return nil, err

}

func (api *API) ShowScore(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	w.Header().Set("Content-Type", "text/html")

	uid, _ := ctx.Value("user_id").(string)

	score, err := api.Interactor.UseCase.GetScore(ctx, uid)
	if err != nil {
		log.Println(err)
	}

	if err := templatego2.TemplateMap["score"].Execute(w, score); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}

}

func (api *API) ShowQuestion(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html")
	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}
	ques, _ := api.Interactor.UseCase.GetQuestionsAnswers(ctx, userId)

	qtemplate := struct {
		Ques []models.QuestionAnswersResponse
	}{
		Ques: ques,
	}

	if err := templatego2.TemplateMap["questions"].Execute(w, qtemplate); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}

}

func (api *API) GetQuestionAnswers(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()
	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}
	res, err := api.Interactor.UseCase.GetQuestionsAnswers(ctx, userId)

	return res, err

}

func (api *API) InsertQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}
	user, err := api.Interactor.UseCase.GetUser(ctx, userId)
	if err != nil {
		//return nil, err
	}
	if user.UserType != constant.UserRoleAdmin {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}

	qCount, err := api.Interactor.UseCase.GetQuestionsCount(ctx)
	if err != nil {
		log.Println(err)
	}

	var qCountTotal, qAnswered int64

	for _, qc := range qCount {
		if qc.AnsweredByUsers >= int64(constant.UserCount) {
			qAnswered = qAnswered + qc.Count
		}
		qCountTotal = qCountTotal + qc.Count
	}
	_, err = api.InsertQuestionAndUploadFile(ctx, r)
	templateData := struct {
		Message       string
		QuestionCount int64
		AnsweredCount int64
	}{}
	if err == nil {
		templateData.Message = "Upload Successful, Insert next Question"
	}
	templateData.QuestionCount = qCountTotal
	templateData.AnsweredCount = qAnswered
	w.Header().Set("Content-Type", "text/html")

	if err := templatego2.TemplateMap["upload"].Execute(w, templateData); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)
	}

}

func (api *API) CreateQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}
	user, err := api.Interactor.UseCase.GetUser(ctx, userId)
	if err != nil {
		//return nil, err
	}
	if user.UserType != constant.UserRoleAdmin {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}

}

// This function returns the filenames(to save in database) of the saved file
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

	ques := r.FormValue("question_text")
	imageId, err := api.FileUpload(ctx, fhsQues[0])
	qid, err := api.Interactor.UseCase.InsertQuestion(ctx, ques, "./upload/"+imageId)
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

func (api *API) InsertQuestionRandom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}
	user, err := api.Interactor.UseCase.GetUser(ctx, userId)
	if err != nil {
		//return nil, err
	}
	if user.UserType != constant.UserRoleAdmin {
		w.WriteHeader(http.StatusUnauthorized)
		//return nil, ErrUnAuthorized
	}
	_, err = api.InsertQuestionAndUploadFileRandom(ctx, r)
	templateData := struct {
		Message string
	}{}
	if err == nil {
		templateData.Message = "Upload Successful, Insert next Question"
	}
	w.Header().Set("Content-Type", "text/html")

	if err := templatego2.TemplateMap["upload_random"].Execute(w, templateData); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)
	}

}

func (api *API) InsertQuestionAndUploadFileRandom(ctx context.Context, r *http.Request) ([]string, error) {

	var fileNames []string
	err := r.ParseMultipartForm(64 << 20) // max size used by FormFile
	if err != nil {
		return fileNames, err
	}
	fhsQues := r.MultipartForm.File["ques_files"]

	var i int
	for _, fh := range fhsQues {

		ques := r.FormValue("question_text")
		imageId, err := api.FileUpload(ctx, fh)
		qid, err := api.Interactor.UseCase.InsertQuestion(ctx, ques, "./upload/"+imageId)
		if err != nil {
			log.Println(err)
		}

		fhsAns := r.MultipartForm.File["ans_files"]

		le := len(fhsAns)
		for j := 0; j < 5 && i < le; j++ {
			imageId, err := api.FileUpload(ctx, fhsAns[i])
			_, err = api.Interactor.UseCase.InsertAns(ctx, qid, "./upload/"+imageId)
			if err != nil {
				log.Println(err)
				continue
			}

			fileNames = append(fileNames, fhsAns[i].Filename)
			i++
		}
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

func (api *API) SubmitAns(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	uid, _ := ctx.Value("user_id").(string)
	if err := r.ParseForm(); err != nil {
		// handle error
	}

	var qid, aid int64

	for key, values := range r.PostForm {
		ques := strings.Split(key, "ques-")
		if len(ques) > 1 {
			qid, _ = strconv.ParseInt(ques[1], 10, 64)
		}
		if len(values) < 1 {
			continue
		}
		ans := strings.Split(values[0], "radio-")

		if len(ans) > 1 {
			aid, _ = strconv.ParseInt(ans[1], 10, 64)
		}
		err := api.Interactor.UseCase.SaveAns(ctx, qid, aid, uid)
		if err != nil {
			log.Println("ERR=", err)
		}
		http.Redirect(w, r, "/score", http.StatusSeeOther)
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
