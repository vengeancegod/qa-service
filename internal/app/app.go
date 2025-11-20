package app

import (
	"context"
	"log"
	"net/http"
	"qa-service/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	server          *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	router := a.setupRouter()
	a.server = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: router,
	}
	return nil
}

func (a *App) setupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// questions routes
	mux.HandleFunc("GET /questions/", a.serviceProvider.QuestionHandler().GetAllQuestions)
	mux.HandleFunc("POST /questions/", a.serviceProvider.QuestionHandler().CreateQuestion)
	mux.HandleFunc("GET /questions/{id}/", a.serviceProvider.QuestionHandler().GetQuestionByID)
	mux.HandleFunc("DELETE /questions/{id}/", a.serviceProvider.QuestionHandler().DeleteQuestionByID)

	// answers routes
	mux.HandleFunc("POST /questions/{id}/answers/", a.serviceProvider.AnswerHandler().AddAnswerByQuestionID)
	mux.HandleFunc("GET /answers/{id}/", a.serviceProvider.AnswerHandler().GetAnswerByID)
	mux.HandleFunc("DELETE /answers/{id}/", a.serviceProvider.AnswerHandler().DeleteAnswerByID)

	return mux
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())

	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	log.Println("Shutting down server")
	return a.server.Shutdown(ctx)
}
