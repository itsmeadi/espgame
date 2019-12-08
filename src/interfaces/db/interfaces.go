package db

import (
	"context"
	models2 "github.com/ESPGame/src/entities/models"
)

type UserDb interface {
	InsertUser(ctx context.Context, user models2.User) (int64, error)
	GetUserByIdPass(ctx context.Context, id, pass string) (models2.User, error)
	GetUserById(ctx context.Context, id string) (models2.User, error)
}
type QuestionDb interface {
	InsertQuestion(ctx context.Context, ques models2.Question) (int64, error)
	GetQuestions(ctx context.Context, answered, limit int64) ([]models2.Question, error)
}
type AnsDb interface {
	InsertAnswer(ctx context.Context, ans models2.Answers) (int64, error)
	GetAnswer(ctx context.Context, qid int64) ([]models2.Answers, error)
	GetOptionsForQuestions(ctx context.Context, qid []int64) ([]models2.Answers, error)
}
type QADb interface {
	InsertUserQuestionAnswer(ctx context.Context, qa models2.QuestionAnswer) (int64, error)
	GetUserQuestionsFromUser(ctx context.Context, uid string, size int64) ([]models2.QuestionAnswer, error)
	GetUserQuestionsFromQues(ctx context.Context, qid, size int64) ([]models2.QuestionAnswer, error)
	UpdateUserQuestionsAnswered(ctx context.Context, correctness int, qid, aid int64, uid string) (int64, error)
}
