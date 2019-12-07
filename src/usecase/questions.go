package usecase

import (
	"context"
	constant2 "github.com/ESPGame/src/entities/constant"
	models2 "github.com/ESPGame/src/entities/models"
	db2 "github.com/ESPGame/src/interfaces/db"
)

type QuestionsUseCase struct {
	UserRepo db2.UserDb
	QuesRepo db2.QuestionDb
	AnsRepo  db2.AnsDb
	QARepo   db2.QADb
}

var _ UseCaseInterface = &QuestionsUseCase{}

type UseCaseInterface interface {
	GetQuestions(ctx context.Context) ([]models2.Question, error)
	InsertQuestion(ctx context.Context, quesText, media string) (int64, error)
	//create question
	GetQuestionsAnswers(ctx context.Context) ([]models2.QuestionAnswersResponse, error)
	SaveAns(ctx context.Context, qID, aID int64, uID string) error
	InsertAns(ctx context.Context, qID int64, media string) (int64, error)
	SignIn(ctx context.Context, userId, pass string) (models2.User, error)
	SignUp(ctx context.Context, user models2.User) error
	GetScore(ctx context.Context, uid string) (models2.Score, error)
}

func (ques *QuestionsUseCase) GetScore(ctx context.Context, uid string) (models2.Score, error) {

	var score models2.Score
	userQ, err := ques.QARepo.GetUserQuestionsFromUser(ctx, uid, 1000)
	if err != nil {
		return score, err
	}
	var correct, pending, incorrect int64
	for _, qa := range userQ {
		switch int(qa.Correctness) {
		case constant2.CorrectnessAnsCorrect:
			correct++
		case constant2.CorrectnessAnsInCorrect:
			incorrect++
		case constant2.CorrectnessAnsPending:
			pending++
		}
	}
	score.TotalAnswer = int64(len(userQ))
	score.PendingAns = pending
	score.CorrectAnswer = correct
	score.PendingAns = pending
	return score, nil
}

func (ques *QuestionsUseCase) GetQuestions(ctx context.Context) ([]models2.Question, error) {
	questions, err := ques.QuesRepo.GetQuestions(ctx, 2, 5) //TODO //get question where answerd status is less than 2
	if err != nil {
		return nil, err
	}
	return questions, nil
}
func (ques *QuestionsUseCase) GetQuestionsAnswers(ctx context.Context) ([]models2.QuestionAnswersResponse, error) {

	q, err := ques.GetQuestions(ctx)
	if err != nil {
		return nil, err
	}
	qIds := make([]int64, 0)
	for _, qu := range q {
		qIds = append(qIds, qu.Id)
	}

	ans, err := ques.AnsRepo.GetOptionsForQuestions(ctx, qIds)
	if err != nil {
		return nil, err
	}
	qaMap := make(map[int64][]models2.Answers)

	for _, ans := range ans {
		qaMap[ans.QuestionId] = append(qaMap[ans.QuestionId], ans)
	}
	var QA []models2.QuestionAnswersResponse

	for _, qu := range q {
		QA = append(QA, models2.QuestionAnswersResponse{
			Ques: qu,
			Ans:  qaMap[qu.Id],
		})
	}
	return QA, nil
}

func (ques *QuestionsUseCase) SaveAns(ctx context.Context, qID, aID int64, uID string) error {
	_, err := ques.QARepo.InsertUserQuestionAnswer(ctx,
		models2.QuestionAnswer{
			UserId:     uID,
			AnswerId:   aID,
			QuestionId: qID,
		})
	return err
}
func (ques *QuestionsUseCase) InsertRandomAns(ctx context.Context, qID int64, media string) (int64, error) {
	return 0, nil
}

func (ques *QuestionsUseCase) InsertAns(ctx context.Context, qID int64, media string) (int64, error) {
	aid, err := ques.AnsRepo.InsertAnswer(ctx,
		models2.Answers{
			QuestionId: qID,
			AnswerText: "",
			MediaUrl:   media,
		})
	return aid, err
}

func (ques *QuestionsUseCase) InsertQuestion(ctx context.Context, quesText, media string) (int64, error) {

	var question models2.Question
	question.Question = quesText
	question.MediaUrl = media

	qid, err := ques.QuesRepo.InsertQuestion(ctx, question)
	return qid, err
}

func (ques *QuestionsUseCase) SignIn(ctx context.Context, userId, pass string) (models2.User, error) {

	user, err := ques.UserRepo.GetUser(ctx, userId, pass)
	if err != nil {
		return user, err
	}
	return user, err
}

func (ques *QuestionsUseCase) SignUp(ctx context.Context, user models2.User) error {

	_, err := ques.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return err
	}
	return err
}
