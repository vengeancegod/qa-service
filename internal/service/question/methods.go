package question

import (
	"errors"
	"qa-service/internal/model"
)

func (s *Service) GetAllQuestions() ([]model.Question, error) {
	question, err := s.questionRepository.GetAllQuestions()
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *Service) CreateQuestion(question *model.Question) error {
	if question.Text == "" {
		return errors.New(model.ErrEmptyQuestion)
	}
	err := s.questionRepository.CreateQuestion(question)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetQuestionByID(id int64) (*model.Question, error) {
	question, err := s.questionRepository.GetQuestionByID(id)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *Service) DeleteQuestionByID(id int64) error {
	err := s.questionRepository.DeleteQuestionByID(id)
	if err != nil {
		return err
	}
	return nil
}
