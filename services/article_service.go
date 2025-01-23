package services

import (
	"myapi/models"
	"myapi/repositories"
)

// AirticleDetailHandler で使うことを想定したサービス
func GetArticleService(airticleID int) (models.Article, error) {
	// TODO : sql.DB 型を手に入れて、変数 db に代入する

	// 1. repositories 層の SelectArticleDetail 関数を呼び出し、記事の詳細データを取得する
	article, err := repositories.SelectArticleDetail(db, airticleID)
	if err != nil {
		return models.Article{}, err
	}

	// 2. コメント一覧を取得する
	commentList, err := repositories.SelectCommentList(db, airticleID)
	if err != nil {
		return models.Article{}, err
	}

	// 3. 取得したコメント一覧を記事データに追加する
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}
