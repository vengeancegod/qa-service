package question

import (
	"errors"
	"qa-service/internal/model"
	rep "qa-service/internal/repository"

	"gorm.io/gorm"
)

var _ rep.QuestionRepository = (*repository)(nil)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) (*repository, error) {
	return &repository{
		DB: db,
	}, nil
}

func (repo *repository) GetAllQuestions() ([]model.Question, error) {
	var questions []model.Question

	if err := repo.DB.Preload("Answers").Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (repo *repository) CreateQuestion(question *model.Question) error {
	return repo.DB.Create(question).Error
}

func (repo *repository) GetQuestionByID(id int64) (*model.Question, error) {
	var question model.Question

	if err := repo.DB.Preload("Answers").First(&question, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(model.QuestionNotFound)
		}
		return nil, err
	}
	return &question, nil
}

func (repo *repository) DeleteQuestionByID(id int64) error {
	result := repo.DB.Delete(&model.Question{}, id)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New(model.QuestionNotFound)
	}

	return nil
}
