package db

import (
	"context"
	models "github.com/ESPGame/src/entities/models"
)

type UserDb interface {
	InsertUser(ctx context.Context, user models.User) (int64, error)
	GetUserByIdPass(ctx context.Context, id, pass string) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.User, error)
}

type QuestionDb interface {
	InsertQuestion(ctx context.Context, ques models.Question) (int64, error)
	GetQuestions(ctx context.Context, uid string, noOfUsersAnswered, limit int) ([]models.Question, error)
	GetQuestionsCount(ctx context.Context) ([]models.QuestionCount, error)
	UpdateUserQuestionsAnsweredCount(ctx context.Context, qid int64) (int64, error)
}

type AnsDb interface {
	InsertAnswer(ctx context.Context, ans models.Answers) (int64, error)
	GetAnswer(ctx context.Context, qid int64) ([]models.Answers, error)
	GetOptionsForQuestions(ctx context.Context, qid []int64) ([]models.Answers, error)
}

type QADb interface {
	InsertUserQuestionAnswer(ctx context.Context, qa models.QuestionAnswer) (int64, error)
	GetUserQuestionsFromUser(ctx context.Context, uid string, size int64) ([]models.QuestionAnswer, error)
	GetUserQuestionsAnsweredCount(ctx context.Context, qid int64) (int64, error)
	GetUserQuestionsFromQues(ctx context.Context, qid, size int64) ([]models.QuestionAnswer, error)
	UpdateUserQuestionsAnswered(ctx context.Context, correctness int, qid, aid int64) (int64, error)
}
