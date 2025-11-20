package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"qa-service/internal/model"
	"qa-service/internal/service"
)

type AnswerHandler struct {
	answerService service.AnswerService
}

func NewAnswerHandler(answerService service.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		answerService: answerService,
	}
}

func (h *AnswerHandler) AddAnswerByQuestionID(w http.ResponseWriter, r *http.Request) {
	questionIDStr := r.PathValue("id")
	questionID, err := strconv.ParseInt(questionIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var answer model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.answerService.AddAnswerByQuestionID(questionID, &answer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func (h *AnswerHandler) GetAnswerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	answer, err := h.answerService.GetAnswerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

func (h *AnswerHandler) DeleteAnswerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	if err := h.answerService.DeleteAnswerByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}