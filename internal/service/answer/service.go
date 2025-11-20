package answer

import "qa-service/internal/repository"

type Service struct {
	answerRepository repository.AnswerRepository
	questionRepository repository.QuestionRepository
}

func NewService(answerRepository repository.AnswerRepository, questionRepository repository.QuestionRepository) *Service {
	return &Service{
		answerRepository: answerRepository,
		questionRepository: questionRepository,
	}
}
