package dto

import (
	"blog_api/model"
)

type ArticleDto struct {
	ID         uint             `gorm:"primary_key"`
	Title      string           `json:"title"`
	Status     int              `json:"status"`
	Created_at model.TimeNormal `json:"created_at"`
}
type ArticleInfoDto struct {
	ID         uint             `gorm:"primary_key"`
	Title      string           `json:"title"`
	Content    string           `json:"content"`
	Status     int              `json:"status"`
	Created_at model.TimeNormal `json:"created_at"`
	Updated_at model.TimeNormal `json:"updated_at"`
}

func ToArticleDto(article model.Article) ArticleDto {
	return ArticleDto{
		ID:         article.ID,
		Title:      article.Title,
		Status:     article.Status,
		Created_at: article.CreatedAt,
	}
}
func ToArticleInfoDto(article model.Article) ArticleInfoDto {
	return ArticleInfoDto{
		ID:         article.ID,
		Title:      article.Title,
		Content:    article.Content,
		Status:     article.Status,
		Created_at: article.CreatedAt,
		Updated_at: article.UpdatedAt,
	}
}

func ToArticleArrayDto(articles []model.Article) (artDTO []ArticleDto) {

	for _, v := range articles {
		artDTO = append(artDTO, ToArticleDto(v))
	}
	return artDTO
}
