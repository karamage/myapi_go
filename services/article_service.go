package services

import (
	"myapi/models"
	"myapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// 1. repositories 層の SelectArticleDetail 関数を呼び出し、記事の詳細データを取得する
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 2. コメント一覧を取得する
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 3. 取得したコメント一覧を記事データに追加する
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	var newArticle models.Article

	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
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

func PostNiceService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// 1. repositories 層の UpdateNiceNum 関数を呼び出し、指定IDの記事のいいね数を+1する
	if err := repositories.UpdateNiceNum(db, articleID); err != nil {
		return models.Article{}, err
	}

	// 2. 更新後の記事データを取得する
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}
