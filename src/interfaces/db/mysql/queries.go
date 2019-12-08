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

	insertQuestion                = "insert into questions (question_text,media_url, answered) values (?,?,?)"
	getQuestionByAnsweredAndLimit = "select id,question_text,media_url, answered from questions where answered<? limit ?"

	insertAnswer           = `insert into answers (question_id, answer_text, media_url) values (?,?,?)`
	getAnswerFromQuestion  = "select id, question_id, answer_text, media_url from answers where question_id=?"
	getOptionsForQuestions = "select id, question_id, answer_text, media_url from answers where question_id in (?)"

	insertUserQuestion           = "insert into user_questions(user_id, question_id,answer_id) values(?,?,?)"
	getUserQuestionsFromUser     = "select id, user_id, question_id,answer_id, correctness from user_questions where user_id=? limit ?"
	getUserQuestionsFromQuestion = "select id, user_id, question_id,answer_id from user_questions where question_id=? limit ?"
	updateUserQuestionsAnswered  = "update user_questions set correctness=? where question_id=? and answer_id=? and user_id<>?"
)

type PreparedQueries struct {
	InsertUser         *sql.Stmt
	GetUserByIdAndPass *sql.Stmt
	GetUserById        *sql.Stmt

	InsertQuestion                *sql.Stmt
	GetQuestionByAnsweredAndLimit *sql.Stmt

	InsertAnswer           *sql.Stmt
	GetAnswerFromQuestion  *sql.Stmt
	GetOptionsForQuestions *sql.Stmt

	InsertUserQuestion           *sql.Stmt
	GetUserQuestionsFromUser     *sql.Stmt
	GetUserQuestionsFromQuestion *sql.Stmt
	UpdateUserQuestionsAnswered  *sql.Stmt
}

func (Conn *DB) initQueries() {
	Conn.queries = &PreparedQueries{
		InsertUser:         Conn.Prepare(insertUser),
		GetUserByIdAndPass: Conn.Prepare(getUserByIdAndPass),
		GetUserById:        Conn.Prepare(getUserById),

		InsertQuestion:                Conn.Prepare(insertQuestion),
		GetQuestionByAnsweredAndLimit: Conn.Prepare(getQuestionByAnsweredAndLimit),
		GetOptionsForQuestions:        Conn.Prepare(getOptionsForQuestions),

		InsertAnswer:          Conn.Prepare(insertAnswer),
		GetAnswerFromQuestion: Conn.Prepare(getAnswerFromQuestion),

		InsertUserQuestion:           Conn.Prepare(insertUserQuestion),
		GetUserQuestionsFromUser:     Conn.Prepare(getUserQuestionsFromUser),
		GetUserQuestionsFromQuestion: Conn.Prepare(getUserQuestionsFromQuestion),
		UpdateUserQuestionsAnswered:  Conn.Prepare(updateUserQuestionsAnswered),
	}

}

func (Conn *DB) Prepare(query string) *sql.Stmt {

	prep, err := Conn.sqlConn.Preparex(query)
	if err != nil {
		log.Fatalf("Cannot Prepare Query=%+v, err=%+v", query, err)
	}
	return prep
}
