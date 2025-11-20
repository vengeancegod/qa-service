package service

import "qa-service/internal/model"

type QuestionService interface {
	GetAllQuestions() ([]model.Question, error)
	CreateQuestion(question *model.Question) error
	GetQuestionByID(id int64) (*model.Question, error) 
	DeleteQuestionByID(id int64) error
}

type AnswerService interface {
	AddAnswerByQuestionID(questionID int64, answer *model.Answer) error 
	GetAnswerByID(id int64) (*model.Answer, error)
	DeleteAnswerByID(id int64) error
}