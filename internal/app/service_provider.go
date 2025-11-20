package app

import (
	"log"
	handler "qa-service/internal/api/handlers"
	"qa-service/internal/config"
	"qa-service/internal/infrastructure/psql"
	"qa-service/internal/repository"
	answerRepository "qa-service/internal/repository/answer"
	questionRepository "qa-service/internal/repository/question"
	"qa-service/internal/service"
	answerService "qa-service/internal/service/answer"
	questionService "qa-service/internal/service/question"

	"gorm.io/gorm"
)

type serviceProvider struct {
	serverConfig       config.HTTPConfig
	dbConfig           config.DBConfig
	db                 *gorm.DB
	questionRepository repository.QuestionRepository
	answerRepository   repository.AnswerRepository
	questionService    service.QuestionService
	answerService      service.AnswerService
	questionHandler    *handler.QuestionHandler
	answerHandler      *handler.AnswerHandler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.serverConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.serverConfig = cfg
	}
	return s.serverConfig
}

func (s *serviceProvider) DBConfig() config.DBConfig {
	if s.dbConfig == nil {
		cfg, err := config.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		s.dbConfig = cfg
	}

	return s.dbConfig
}

func (s *serviceProvider) DB() *gorm.DB {
	if s.db == nil {
		db, err := psql.InitDB(s.DBConfig())
		if err != nil {
			log.Fatalf("failed to connect to database: %s", err.Error())
		}

		s.db = db
	}

	return s.db
}

func (s *serviceProvider) QuestionRepository() repository.QuestionRepository {
	if s.questionRepository == nil {
		repo, err := questionRepository.NewRepository(s.DB())
		if err != nil {
			log.Fatalf("failed to create question repository: %s", err.Error())
		}
		s.questionRepository = repo
	}

	return s.questionRepository
}

func (s *serviceProvider) AnswerRepository() repository.AnswerRepository {
	if s.answerRepository == nil {
		repo, err := answerRepository.NewRepository(s.DB())
		if err != nil {
			log.Fatalf("failed to create answer repository: %s", err.Error())
		}
		s.answerRepository = repo
	}

	return s.answerRepository
}

func (s *serviceProvider) QuestionService() service.QuestionService {
	if s.questionService == nil {
		s.questionService = questionService.NewService(
			s.QuestionRepository(),
		)
	}

	return s.questionService
}

func (s *serviceProvider) AnswerService() service.AnswerService {
	if s.answerService == nil {
		s.answerService = answerService.NewService(
			s.AnswerRepository(),
			s.QuestionRepository(),
		)
	}

	return s.answerService
}

func (s *serviceProvider) QuestionHandler() *handler.QuestionHandler {
	if s.questionHandler == nil {
		s.questionHandler = handler.NewQuestionHandler(s.QuestionService())
	}

	return s.questionHandler
}

func (s *serviceProvider) AnswerHandler() *handler.AnswerHandler {
	if s.answerHandler == nil {
		s.answerHandler = handler.NewAnswerHandler(s.AnswerService())
	}

	return s.answerHandler
}
