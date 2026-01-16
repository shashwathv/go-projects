package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"calc_service/internal/middleware"
	"calc_service/internal/models"
	"calc_service/internal/service"
	"calc_service/internal/storage"
	"calc_service/internal/validation"
)

type Handler struct {
	repo storage.Repository
}

func New(repo storage.Repository) *Handler {
	return &Handler{repo: repo}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, ErrorResponse{
		Error: err.Error(),
	})
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var req TwoNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	a, b, err := validation.ValidateTwoNumbers(req.Number1, req.Number2)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	result := service.Add(a, b)

	calc := models.Calculation{
		Operation: "add",
		Operands:  []int{a, b},
		Result:    result,
		RequestID: middleware.GetRequestID(r.Context()),
		UserID:    "anonymous",
		CreatedAt: time.Now(),
	}

	if err := h.repo.SaveCalculation(r.Context(), calc); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, ResultResponse{Result: result})
}

func (h *Handler) Subtract(w http.ResponseWriter, r *http.Request) {
	var req TwoNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	a, b, err := validation.ValidateTwoNumbers(req.Number1, req.Number2)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	result := service.Subtract(a, b)

	calc := models.Calculation{
		Operation: "subtract",
		Operands:  []int{a, b},
		Result:    result,
		RequestID: middleware.GetRequestID(r.Context()),
		UserID:    "anonymous",
		CreatedAt: time.Now(),
	}

	if err := h.repo.SaveCalculation(r.Context(), calc); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, ResultResponse{Result: result})
}

func (h *Handler) Multiply(w http.ResponseWriter, r *http.Request) {
	var req TwoNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	a, b, err := validation.ValidateTwoNumbers(req.Number1, req.Number2)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	result := service.Multiply(a, b)

	calc := models.Calculation{
		Operation: "multiply",
		Operands:  []int{a, b},
		Result:    result,
		RequestID: middleware.GetRequestID(r.Context()),
		UserID:    "anonymous",
		CreatedAt: time.Now(),
	}

	if err := h.repo.SaveCalculation(r.Context(), calc); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, ResultResponse{Result: result})
}

func (h *Handler) Divide(w http.ResponseWriter, r *http.Request) {
	var req DivideRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	a, b, err := validation.ValidateDivide(req.Dividend, req.Divisor)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	result, err := service.Divide(a, b)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	calc := models.Calculation{
		Operation: "divide",
		Operands:  []int{a, b},
		Result:    result,
		RequestID: middleware.GetRequestID(r.Context()),
		UserID:    "anonymous",
		CreatedAt: time.Now(),
	}

	if err := h.repo.SaveCalculation(r.Context(), calc); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, ResultResponse{Result: result})
}

func (h *Handler) Sum(w http.ResponseWriter, r *http.Request) {
	var req SumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	numbers, err := validation.ValidateSum(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	result, err := service.Sum(numbers)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	calc := models.Calculation{
		Operation: "sum",
		Operands:  numbers,
		Result:    result,
		RequestID: middleware.GetRequestID(r.Context()),
		UserID:    "anonymous",
		CreatedAt: time.Now(),
	}

	if err := h.repo.SaveCalculation(r.Context(), calc); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, ResultResponse{Result: result})
}
