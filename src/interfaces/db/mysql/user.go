package mysql

import (
	"context"
	"database/sql"
	models2 "github.com/ESPGame/src/entities/models"
	"github.com/jmoiron/sqlx"
)

func (Conn *DB) InsertUser(ctx context.Context, user models2.User) (int64, error) {

	res, err := Conn.queries.InsertUser.Exec(user.Id, user.Name, user.Password, user.UserType)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetUser(ctx context.Context, id, pass string) (models2.User, error) {

	var user models2.User

	err := Conn.queries.GetUserByIdAndPass.GetContext(ctx, &user, id, pass)
	if err == sql.ErrNoRows {
		return user, nil
	}
	return user, err
}

func (Conn *DB) InsertQuestion(ctx context.Context, ques models2.Question) (int64, error) {

	res, err := Conn.queries.InsertQuestion.Exec(ques.Question, ques.MediaUrl, ques.Answered)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}
func (Conn *DB) GetQuestions(ctx context.Context, answered, limit int64) ([]models2.Question, error) {

	var ques []models2.Question
	err := Conn.queries.GetQuestionByAnsweredAndLimit.SelectContext(ctx, &ques, answered, limit)
	return ques, err
}

func (Conn *DB) InsertAnswer(ctx context.Context, ans models2.Answers) (int64, error) {

	res, err := Conn.queries.InsertAnswer.Exec(ans.QuestionId, ans.AnswerText, ans.MediaUrl)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetAnswer(ctx context.Context, qid int64) ([]models2.Answers, error) {

	var ans []models2.Answers

	err := Conn.queries.GetAnswerFromQuestion.GetContext(ctx, &ans, qid)
	return ans, err
}

func (Conn *DB) GetOptionsForQuestions(ctx context.Context, qid []int64) ([]models2.Answers, error) {

	var ans []models2.Answers

	query, args, err := sqlx.In(getOptionsForQuestions, qid)

	if err != nil {
		return nil, err
	}

	query = Conn.sqlConn.Rebind(query)
	err = Conn.sqlConn.Select(&ans, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return ans, err
}

func (Conn *DB) InsertUserQuestionAnswer(ctx context.Context, qa models2.QuestionAnswer) (int64, error) {

	res, err := Conn.queries.InsertUserQuestion.Exec(qa.UserId, qa.QuestionId, qa.AnswerId)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetUserQuestionsFromUser(ctx context.Context, uid string, size int64) ([]models2.QuestionAnswer, error) {

	var qa []models2.QuestionAnswer

	err := Conn.queries.GetUserQuestionsFromUser.SelectContext(ctx, &qa, uid, size)
	return qa, err
}

func (Conn *DB) GetUserQuestionsFromQues(ctx context.Context, qid, size int64) ([]models2.QuestionAnswer, error) {

	var qa []models2.QuestionAnswer

	err := Conn.queries.GetUserQuestionsFromQuestion.GetContext(ctx, &qa, qid, size)
	return qa, err
}
