package services

import (
	"myapi/models"
	"myapi/repositories"
)

// AirticleDetailHandler で使うことを想定したサービス
func GetArticleService(airticleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

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

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	var newArticle models.Article

	// 1. repositories 層の InsertArticle 関数を呼び出し、記事データを保存する
	if newArticle, err = repositories.InsertArticle(db, article); err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// repositories 層の SelectArticleList 関数を呼び出し、指定ページの記事一覧を取得する
	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}
