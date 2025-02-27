package controllers

import (
	"encoding/json"
	"myapi/controllers/services"
	"myapi/models"
	"net/http"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(service services.CommentServicer) *CommentController {
	return &CommentController{service: service}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid comment data", http.StatusBadRequest)
		return
	}

	newComment, err := c.service.PostCommentService(comment)
	if err != nil {
		http.Error(w, "Failed to post comment", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newComment)
}
