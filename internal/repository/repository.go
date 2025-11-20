package repository

import "qa-service/internal/model"

type QuestionRepository interface {
	GetAllQuestions() ([]model.Question, error)
	CreateQuestion(question *model.Question) error
	GetQuestionByID(id int64) (*model.Question, error)
	DeleteQuestionByID(id int64) error
}

type AnswerRepository interface {
	AddAnswerByQuestionID(questionID int64, answer *model.Answer) error
	GetAnswerByID(id int64) (*model.Answer, error) //vopros
	DeleteAnswerByID(id int64) error
}
