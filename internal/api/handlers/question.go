package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"qa-service/internal/model"
	"qa-service/internal/service"
)

type QuestionHandler struct {
	questionService service.QuestionService
}

func NewQuestionHandler(questionService service.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.questionService.GetAllQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question model.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.questionService.CreateQuestion(&question); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func (h *QuestionHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	question, err := h.questionService.GetQuestionByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func (h *QuestionHandler) DeleteQuestionByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	if err := h.questionService.DeleteQuestionByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
