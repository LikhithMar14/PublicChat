package main

import (
	"net/http"

	"github.com/LikhithMar14/social/internal/models"
)

type CreatePostPayload struct{
	Content string `json:"content"`
	Title string `json:"title"`
	Tags []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	var payload CreatePostPayload
	err := readJSON(w, r, &payload)
	if err != nil{
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	post := &models.Post{
		Content: payload.Content,
		Title: payload.Title,
		UserID: 1,
		Tags: payload.Tags,
	}
	err = app.store.Posts.Create(ctx, post)
	if err != nil{
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
