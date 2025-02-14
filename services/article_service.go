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

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	// repositories 層の SelectArticleList 関数を呼び出し、指定ページの記事一覧を取得する
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	// 1. repositories 層の UpdateNiceNum 関数を呼び出し、指定IDの記事のいいね数を+1する
	if err := repositories.UpdateNiceNum(s.db, article.ID); err != nil {
		return models.Article{}, err
	}

	// 2. 更新後の記事データを取得する
	updatedArticle, err := repositories.SelectArticleDetail(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return updatedArticle, nil
}
