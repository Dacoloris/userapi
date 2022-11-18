package handler

import (
	"net/http"
	"refactoring/internal/handler/requests"
	"refactoring/internal/handler/responses"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.useCase.GetUsers(r.Context())
	if err != nil {
		_ = render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, users)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := requests.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	id, err := h.useCase.CreateUser(r.Context(), request)
	if err != nil {
		_ = render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.useCase.GetUser(r.Context(), id)
	if err != nil {
		_ = render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}

func (h *Handler) UpdateUserName(w http.ResponseWriter, r *http.Request) {
	request := requests.UpdateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	if err := h.useCase.UpdateUserName(r.Context(), id, request.DisplayName); err != nil {
		_ = render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.useCase.DeleteUser(r.Context(), id); err != nil {
		_ = render.Render(w, r, responses.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
