package answer

import (
	"errors"
	"qa-service/internal/model"
	"qa-service/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAnswerRepository struct {
	mock.Mock
}

func (m *MockAnswerRepository) AddAnswerByQuestionID(questionID int64, answer *model.Answer) error {
	args := m.Called(questionID, answer)
	return args.Error(0)
}

func (m *MockAnswerRepository) GetAnswerByID(id int64) (*model.Answer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Answer), args.Error(1)
}

func (m *MockAnswerRepository) DeleteAnswerByID(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) GetAllQuestions() ([]model.Question, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Question), args.Error(1)
}

func (m *MockQuestionRepository) CreateQuestion(question *model.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionRepository) GetQuestionByID(id int64) (*model.Question, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Question), args.Error(1)
}

func (m *MockQuestionRepository) DeleteQuestionByID(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

var _ repository.QuestionRepository = (*MockQuestionRepository)(nil)

func TestAddAnswerByQuestionID_Success(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	questionID := int64(1)
	answer := &model.Answer{
		Text: "Test answer",
	}

	question := &model.Question{
		ID:   questionID,
		Text: "Test question",
	}

	mockQuestionRepo.On("GetQuestionByID", questionID).Return(question, nil)
	mockAnswerRepo.On("AddAnswerByQuestionID", questionID, answer).Return(nil)

	err := service.AddAnswerByQuestionID(questionID, answer)

	assert.NoError(t, err)
	assert.NotNil(t, answer.UserID)
	assert.Equal(t, questionID, answer.QuestionID)
	mockQuestionRepo.AssertExpectations(t)
	mockAnswerRepo.AssertExpectations(t)
}

func TestAddAnswerByQuestionID_InvalidQuestionID(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	testCases := []struct {
		name       string
		questionID int64
	}{
		{"Zero ID", 0},
		{"Negative ID", -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			answer := &model.Answer{Text: "Test answer"}
			err := service.AddAnswerByQuestionID(tc.questionID, answer)

			assert.Error(t, err)
			assert.Equal(t, model.ErrInvalidQuestionID, err.Error())
		})
	}
}

func TestAddAnswerByQuestionID_EmptyText(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	answer := &model.Answer{
		Text: "",
	}

	err := service.AddAnswerByQuestionID(1, answer)

	assert.Error(t, err)
	assert.Equal(t, model.ErrEmptyAnsText, err.Error())
}

func TestAddAnswerByQuestionID_QuestionNotFound(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	questionID := int64(999)
	answer := &model.Answer{
		Text: "Test answer",
	}

	mockQuestionRepo.On("GetQuestionByID", questionID).Return(nil, errors.New("not found"))

	err := service.AddAnswerByQuestionID(questionID, answer)

	assert.Error(t, err)
	assert.Equal(t, model.QuestionNotFound, err.Error())
	mockQuestionRepo.AssertExpectations(t)
}

func TestAddAnswerByQuestionID_RepositoryError(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	questionID := int64(1)
	answer := &model.Answer{
		Text: "Test answer",
	}

	question := &model.Question{
		ID:   questionID,
		Text: "Test question",
	}

	mockQuestionRepo.On("GetQuestionByID", questionID).Return(question, nil)
	mockAnswerRepo.On("AddAnswerByQuestionID", questionID, answer).Return(errors.New("database error"))

	err := service.AddAnswerByQuestionID(questionID, answer)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
}

func TestGetAnswerByID_Success(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	expectedAnswer := &model.Answer{
		ID:   1,
		Text: "Test answer",
	}

	mockAnswerRepo.On("GetAnswerByID", int64(1)).Return(expectedAnswer, nil)

	answer, err := service.GetAnswerByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAnswer, answer)
	mockAnswerRepo.AssertExpectations(t)
}

func TestGetAnswerByID_InvalidID(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	testCases := []struct {
		name string
		id   int64
	}{
		{"Zero ID", 0},
		{"Negative ID", -5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			answer, err := service.GetAnswerByID(tc.id)

			assert.Error(t, err)
			assert.Nil(t, answer)
			assert.Equal(t, model.ErrAnswerID, err.Error())
		})
	}
}

func TestGetAnswerByID_NotFound(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	mockAnswerRepo.On("GetAnswerByID", int64(999)).Return(nil, errors.New("not found"))

	answer, err := service.GetAnswerByID(999)

	assert.Error(t, err)
	assert.Nil(t, answer)
	mockAnswerRepo.AssertExpectations(t)
}

func TestDeleteAnswerByID_Success(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	mockAnswerRepo.On("DeleteAnswerByID", int64(1)).Return(nil)

	err := service.DeleteAnswerByID(1)

	assert.NoError(t, err)
	mockAnswerRepo.AssertExpectations(t)
}

func TestDeleteAnswerByID_Error(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)
	mockQuestionRepo := new(MockQuestionRepository)

	service := &Service{
		answerRepository:   mockAnswerRepo,
		questionRepository: mockQuestionRepo,
	}

	mockAnswerRepo.On("DeleteAnswerByID", int64(1)).Return(errors.New("delete failed"))

	err := service.DeleteAnswerByID(1)

	assert.Error(t, err)
	assert.Equal(t, "delete failed", err.Error())
	mockAnswerRepo.AssertExpectations(t)
}
