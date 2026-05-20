package http

import (
	"encoding/json"
	"net/http"

	"github.com/diazjohan98/go-virtual-queue-system/internal/usecase"
)

type JoinQueueRequest struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
}

type QueueHandler struct {
	enqueueUseCase *usecase.EnqueueUserUseCase
}

func NewQueueHandler(uc *usecase.EnqueueUserUseCase) *QueueHandler {
	return &QueueHandler{
		enqueueUseCase: uc,
	}
}

func (h *QueueHandler) Join(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido, debe ser POST"})
		return
	}

	var req JoinQueueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Solicitud inválida"})
		return
	}

	if req.EventID == "" || req.UserID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Los campos event_id y user_id son obligatorios"})
		return
	}

	err = h.enqueueUseCase.Execute(r.Context(), req.EventID, req.UserID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "¡Te has unido exitosamente a la fila virtual! Espera tu turno.",
	})
}
