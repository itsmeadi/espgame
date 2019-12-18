package usecase

import (
	"context"
	"github.com/ESPGame/src/entities/constant"
	models "github.com/ESPGame/src/entities/models"
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
	GetQuestionsForUser(ctx context.Context, uid string) ([]models.Question, error)
	InsertQuestion(ctx context.Context, quesText, media string) (int64, error)
	//create question
	GetQuestionsAnswers(ctx context.Context, uid string) ([]models.QuestionAnswersResponse, error)
	GetQuestionsCount(ctx context.Context) ([]models.QuestionCount, error)
	SaveAns(ctx context.Context, qID, aID int64, uID string) error
	InsertAns(ctx context.Context, qID int64, media string) (int64, error)
	SignIn(ctx context.Context, userId, pass string) (models.User, error)
	GetUser(ctx context.Context, userId string) (models.User, error)
	SignUp(ctx context.Context, user models.User) error
	GetScore(ctx context.Context, uid string) (models.Score, error)
}

func (ques *QuestionsUseCase) GetScore(ctx context.Context, uid string) (models.Score, error) {

	var score models.Score
	userQ, err := ques.QARepo.GetUserQuestionsFromUser(ctx, uid, 1000)
	if err != nil {
		return score, err
	}
	var correct int64
	for _, qa := range userQ {
		if qa.Correctness == constant.CorrectnessAnsCorrect {
			correct++
		}
	}
	score.TotalAnswer = int64(len(userQ))
	score.CorrectAnswer = correct
	return score, nil
}

func (ques *QuestionsUseCase) SaveAns(ctx context.Context, qID, aID int64, uID string) error {

	correctness := constant.CorrectnessAnsPending

	//set correctness of other user question with same answer to constant.CorrectnessAnsCorrect
	sameAnsByUsers, err := ques.QARepo.UpdateUserQuestionsAnswered(ctx, constant.CorrectnessAnsCorrect, qID, aID)
	if err != nil {
		return err
	}

	if sameAnsByUsers > 0 {
		correctness = constant.CorrectnessAnsCorrect
	}
	_, err = ques.QARepo.InsertUserQuestionAnswer(ctx,
		models.QuestionAnswer{
			UserId:      uID,
			AnswerId:    aID,
			QuestionId:  qID,
			Correctness: correctness,
		})

	//increment question answered by count
	_, err = ques.QuesRepo.UpdateUserQuestionsAnsweredCount(ctx, qID)

	return err
}

func (ques *QuestionsUseCase) GetQuestionsForUser(ctx context.Context, uid string) ([]models.Question, error) {
	questions, err := ques.QuesRepo.GetQuestions(ctx, uid, constant.UserCount, constant.QuestionsLimit)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (ques *QuestionsUseCase) GetQuestionsCount(ctx context.Context) ([]models.QuestionCount, error)  {
	questions, err := ques.QuesRepo.GetQuestionsCount(ctx)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (ques *QuestionsUseCase) GetQuestionsAnswers(ctx context.Context, uid string) ([]models.QuestionAnswersResponse, error) {

	q, err := ques.GetQuestionsForUser(ctx, uid)
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
	qaMap := make(map[int64][]models.Answers)

	for _, ans := range ans {
		qaMap[ans.QuestionId] = append(qaMap[ans.QuestionId], ans)
	}
	var QA []models.QuestionAnswersResponse

	for _, qu := range q {
		QA = append(QA, models.QuestionAnswersResponse{
			Ques: qu,
			Ans:  qaMap[qu.Id],
		})
	}
	return QA, nil
}

func (ques *QuestionsUseCase) InsertRandomAns(ctx context.Context, qID int64, media string) (int64, error) {
	return 0, nil
}

func (ques *QuestionsUseCase) InsertAns(ctx context.Context, qID int64, media string) (int64, error) {
	aid, err := ques.AnsRepo.InsertAnswer(ctx,
		models.Answers{
			QuestionId: qID,
			AnswerText: "",
			MediaUrl:   media,
		})
	return aid, err
}

func (ques *QuestionsUseCase) InsertQuestion(ctx context.Context, quesText, media string) (int64, error) {

	var question models.Question
	question.Question = quesText
	question.MediaUrl = media

	qid, err := ques.QuesRepo.InsertQuestion(ctx, question)
	return qid, err
}

func (ques *QuestionsUseCase) SignIn(ctx context.Context, userId, pass string) (models.User, error) {

	user, err := ques.UserRepo.GetUserByIdPass(ctx, userId, pass)
	if err != nil {
		return user, err
	}
	return user, err
}

func (ques *QuestionsUseCase) GetUser(ctx context.Context, userId string) (models.User, error) {

	user, err := ques.UserRepo.GetUserById(ctx, userId)
	if err != nil {
		return user, err
	}
	return user, err
}

func (ques *QuestionsUseCase) SignUp(ctx context.Context, user models.User) error {

	_, err := ques.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return err
	}
	return err
}
