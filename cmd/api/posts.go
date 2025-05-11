package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LikhithMar14/social/internal/models"
	"github.com/LikhithMar14/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Content string   `json:"content" validate:"required,max=1000"`
	Title   string   `json:"title" validate:"required,max=100"`
	Tags    []string `json:"tags" validate:"required,max=10"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload CreatePostPayload
	err := readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResposne(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResposne(w, r, err)
		return
	}
	post := &models.Post{
		Content: payload.Content,
		Title:   payload.Title,
		UserID:  1,
		Tags:    payload.Tags,
	}
	err = app.store.Posts.Create(ctx, post)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	postID := chi.URLParam(r, "id")
	postIDInt, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		app.badRequestResposne(w, r, err)
		return
	}

	post, err := app.store.Posts.GetByID(ctx, postIDInt)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFound(w, r)
		default:
			app.internalServerError(w, r, err)
		}
	}
	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
