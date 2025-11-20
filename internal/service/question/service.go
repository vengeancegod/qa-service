package question

import "qa-service/internal/repository"

type Service struct {
	questionRepository repository.QuestionRepository
}

func NewService(questionRepository repository.QuestionRepository) *Service {
	return &Service{
		questionRepository: questionRepository,
	}
}
