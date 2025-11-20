package answer

import (
	"errors"
	"qa-service/internal/model"

	"github.com/google/uuid"
)

func (s *Service) AddAnswerByQuestionID(questionID int64, answer *model.Answer) error {
	if questionID <= 0 {
		return errors.New(model.ErrInvalidQuestionID)
	}
	if answer.Text == "" {
		return errors.New(model.ErrEmptyAnsText)
	}
	_, err := s.questionRepository.GetQuestionByID(questionID)
	if err != nil {
		return errors.New(model.QuestionNotFound)
	}
	answer.UserID = uuid.New()
	answer.QuestionID = questionID

	err = s.answerRepository.AddAnswerByQuestionID(questionID, answer)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAnswerByID(id int64) (*model.Answer, error) {
	if id <= 0 {
		return nil, errors.New(model.ErrAnswerID)
	}
	answer, err := s.answerRepository.GetAnswerByID(id)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (s *Service) DeleteAnswerByID(id int64) error {
	err := s.answerRepository.DeleteAnswerByID(id)
	if err != nil {
		return err
	}
	return nil
}
