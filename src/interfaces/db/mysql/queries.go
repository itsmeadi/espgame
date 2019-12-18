package mysql

import (
	sql "github.com/jmoiron/sqlx"

	//"database/sql"
	"log"
)

var (
	insertUser         = "insert into user (id, name,password,usertype) values (?,?,?,?)"
	getUserByIdAndPass = "select id,name,usertype from user where id=? and password=?"
	getUserById        = "select id,name,usertype from user where id=?"

	insertQuestion                = "insert into questions (question_text,media_url) values (?,?)"
	getQuestionByAnsweredAndLimit = `select id, question_text, media_url from ESPGAME.questions where answered_by_users<? and 
									id not in(select question_id from ESPGAME.user_questions_answers where user_id=?) order by answered_by_users desc limit ?`
	getQuestionsCount = `select count(id) as count, answered_by_users from ESPGAME.questions group by answered_by_users`
	updateUserQuestionsAnsweredCount = "update questions set answered_by_users=answered_by_users+1 where id=?"

	insertAnswer           = `insert into answers (question_id, answer_text, media_url) values (?,?,?)`
	getAnswerFromQuestion  = "select id, question_id, answer_text, media_url from answers where question_id=?"
	getOptionsForQuestions = "select id, question_id, answer_text, media_url from answers where question_id in (?)"

	insertUserQuestion            = "insert into user_questions_answers(user_id, question_id,answer_id,correctness) values(?,?,?,?)"
	getUserQuestionsFromUser      = "select id, user_id, question_id,answer_id, correctness from user_questions_answers where user_id=? limit ?"
	getUserQuestionsAnsweredCount = "select count(id) from user_questions_answers where question_id=?"
	getUserQuestionsFromQuestion  = "select id, user_id, question_id,answer_id from user_questions_answers where question_id=? limit ?"
	updateUserQuestionsAnswered   = "update user_questions_answers set correctness=? where question_id=? and answer_id=?"
)

type PreparedQueries struct {
	InsertUser         *sql.Stmt
	GetUserByIdAndPass *sql.Stmt
	GetUserById        *sql.Stmt

	InsertQuestion                *sql.Stmt
	GetQuestionsCount                *sql.Stmt
	GetQuestionByAnsweredAndLimit *sql.Stmt

	InsertAnswer           *sql.Stmt
	GetAnswerFromQuestion  *sql.Stmt
	GetOptionsForQuestions *sql.Stmt

	InsertUserQuestion               *sql.Stmt
	GetUserQuestionsFromUser         *sql.Stmt
	GetUserQuestionsAnsweredCount    *sql.Stmt
	GetUserQuestionsFromQuestion     *sql.Stmt
	UpdateUserQuestionsAnswered      *sql.Stmt
	UpdateUserQuestionsAnsweredCount *sql.Stmt

	CreateGroup *sql.Stmt
}

func (Conn *DB) initQueries() {
	Conn.queries = &PreparedQueries{
		InsertUser:         Conn.Prepare(insertUser),
		GetUserByIdAndPass: Conn.Prepare(getUserByIdAndPass),
		GetUserById:        Conn.Prepare(getUserById),

		InsertQuestion:                Conn.Prepare(insertQuestion),
		GetQuestionsCount:                Conn.Prepare(getQuestionsCount),
		GetQuestionByAnsweredAndLimit: Conn.Prepare(getQuestionByAnsweredAndLimit),
		GetOptionsForQuestions:        Conn.Prepare(getOptionsForQuestions),

		InsertAnswer:          Conn.Prepare(insertAnswer),
		GetAnswerFromQuestion: Conn.Prepare(getAnswerFromQuestion),

		InsertUserQuestion:               Conn.Prepare(insertUserQuestion),
		GetUserQuestionsFromUser:         Conn.Prepare(getUserQuestionsFromUser),
		GetUserQuestionsAnsweredCount:    Conn.Prepare(getUserQuestionsAnsweredCount),
		GetUserQuestionsFromQuestion:     Conn.Prepare(getUserQuestionsFromQuestion),
		UpdateUserQuestionsAnswered:      Conn.Prepare(updateUserQuestionsAnswered),
		UpdateUserQuestionsAnsweredCount: Conn.Prepare(updateUserQuestionsAnsweredCount),
	}

}

func (Conn *DB) Prepare(query string) *sql.Stmt {

	prep, err := Conn.sqlConn.Preparex(query)
	if err != nil {
		log.Fatalf("Cannot Prepare Query=%+v, err=%+v", query, err)
	}
	return prep
}
