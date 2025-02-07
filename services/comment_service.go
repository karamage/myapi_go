package services

import (
	"myapi/models"
	"myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	// repositories 層の InsertComment 関数を呼び出し、新しいコメントをデータベースに保存する
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
