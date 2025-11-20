package answer

import (
	"errors"
	"qa-service/internal/model"
	rep "qa-service/internal/repository"

	"gorm.io/gorm"
)

var _ rep.AnswerRepository = (*repository)(nil)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) (*repository, error) {
	return &repository{
		DB: db,
	}, nil
}

func (repo *repository) AddAnswerByQuestionID(questionID int64, answer *model.Answer) error {
	var question model.Question
	if err := repo.DB.First(&question, questionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(model.QuestionNotFound)
		}
		return err
	}
	answer.QuestionID = questionID

	if err := repo.DB.Create(answer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetAnswerByID(id int64) (*model.Answer, error) {
	var answer model.Answer
	if err := repo.DB.First(&answer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(model.AnswerNotFound)
		}
		return nil, err
	}
	return &answer, nil
}

func (repo *repository) DeleteAnswerByID(id int64) error {
	result := repo.DB.Delete(&model.Answer{}, id)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New(model.AnswerNotFound)
	}

	return nil
}
