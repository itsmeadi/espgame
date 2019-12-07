package models

type User struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	UserType string `db:"usertype"`
}

type Question struct {
	Id       int64  `db:"id"`
	Question string `db:"question_text"`
	MediaUrl string `db:"media_url"`
	Answered int    `db:"answered"`
}

type Answers struct {
	Id         int64  `db:"id"`
	QuestionId int64  `db:"question_id"`
	AnswerText string `db:"answer_text"`
	MediaUrl   string `db:"media_url"`
	Answered   int    `db:"answered"`
}

type QuestionAnswersResponse struct {
	Ques Question
	Ans  []Answers
}

type QuestionAnswer struct {
	Id          int64  `db:"id"`
	UserId      string `db:"user_id"`
	QuestionId  int64  `db:"question_id"`
	AnswerId    int64  `db:"answer_id"`
	Correctness int64  `db:"correctness"`
}

type Score struct {
	TotalAnswer   int64 `json:"total_answer"`
	CorrectAnswer int64 `json:"correct_answer"`
	PendingAns    int64 `json:"pending_ans"`
	IncorrectAns  int64 `json:"incorrect_ans"`
}
