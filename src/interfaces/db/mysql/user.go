package mysql

import (
	"context"
	"database/sql"
	"github.com/ESPGame/src/entities/models"
	"github.com/ESPGame/src/interfaces/db"
	"github.com/jmoiron/sqlx"
)

var _ db.UserDb = &DB{}

func (Conn *DB) InsertUser(ctx context.Context, user models.User) (int64, error) {

	res, err := Conn.queries.InsertUser.Exec(user.Id, user.Name, user.Password, user.UserType)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetUserByIdPass(ctx context.Context, id, pass string) (models.User, error) {

	var user models.User

	err := Conn.queries.GetUserByIdAndPass.GetContext(ctx, &user, id, pass)
	if err == sql.ErrNoRows {
		//return user, nil
	}
	return user, err
}

func (Conn *DB) GetUserById(ctx context.Context, id string) (models.User, error) {

	var user models.User

	err := Conn.queries.GetUserById.GetContext(ctx, &user, id)
	if err == sql.ErrNoRows {
		//return user, nil
	}
	return user, err
}

func (Conn *DB) InsertQuestion(ctx context.Context, ques models.Question) (int64, error) {

	res, err := Conn.queries.InsertQuestion.Exec(ques.Question, ques.MediaUrl)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetQuestions(ctx context.Context, uid string, noOfUserAnswered int, limit int) ([]models.Question, error) {

	var ques []models.Question
	err := Conn.queries.GetQuestionByAnsweredAndLimit.SelectContext(ctx, &ques, noOfUserAnswered, uid, limit)
	return ques, err
}

func (Conn *DB) InsertAnswer(ctx context.Context, ans models.Answers) (int64, error) {

	res, err := Conn.queries.InsertAnswer.Exec(ans.QuestionId, ans.AnswerText, ans.MediaUrl)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetAnswer(ctx context.Context, qid int64) ([]models.Answers, error) {

	var ans []models.Answers

	err := Conn.queries.GetAnswerFromQuestion.GetContext(ctx, &ans, qid)
	return ans, err
}
func (Conn *DB) GetQuestionsCount(ctx context.Context) ([]models.QuestionCount, error) {

	var qCount []models.QuestionCount

	err := Conn.queries.GetQuestionsCount.SelectContext(ctx, &qCount)
	return qCount, err
}

func (Conn *DB) GetOptionsForQuestions(ctx context.Context, qid []int64) ([]models.Answers, error) {

	var ans []models.Answers

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

func (Conn *DB) InsertUserQuestionAnswer(ctx context.Context, qa models.QuestionAnswer) (int64, error) {

	res, err := Conn.queries.InsertUserQuestion.Exec(qa.UserId, qa.QuestionId, qa.AnswerId, qa.Correctness)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) UpdateUserQuestionsAnswered(ctx context.Context, correctness int, qid, aid int64) (int64, error) {

	res, err := Conn.queries.UpdateUserQuestionsAnswered.Exec(correctness, qid, aid)
	if err != nil {
		return 0, err
	}
	rowAff, _ := res.RowsAffected()
	return rowAff, err
}

func (Conn *DB) UpdateUserQuestionsAnsweredCount(ctx context.Context, qid int64) (int64, error) {

	res, err := Conn.queries.UpdateUserQuestionsAnsweredCount.Exec(qid)
	if err != nil {
		return 0, err
	}
	rowAff, _ := res.RowsAffected()
	return rowAff, err
}

func (Conn *DB) GetUserQuestionsFromUser(ctx context.Context, uid string, size int64) ([]models.QuestionAnswer, error) {

	var qa []models.QuestionAnswer

	err := Conn.queries.GetUserQuestionsFromUser.SelectContext(ctx, &qa, uid, size)
	return qa, err
}
func (Conn *DB) GetUserQuestionsAnsweredCount(ctx context.Context, qid int64) (int64, error) {

	var count []int64
	var ansCount int64

	err := Conn.queries.GetUserQuestionsAnsweredCount.SelectContext(ctx, &count, qid)
	if len(count) == 1 {
		ansCount = count[0]
	}
	return ansCount, err
}

func (Conn *DB) GetUserQuestionsFromQues(ctx context.Context, qid, size int64) ([]models.QuestionAnswer, error) {

	var qa []models.QuestionAnswer

	err := Conn.queries.GetUserQuestionsFromQuestion.GetContext(ctx, &qa, qid, size)
	return qa, err
}
