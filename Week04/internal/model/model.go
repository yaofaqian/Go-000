package model

//文章内容表
type ArticleContent struct {
	ArticleId          int64
	ArticleTitle       string
	ArticleAuthor      string
	ArticlePublishTime string
	ArticleHits        int64
	ArticleContent     string
	UserId             int64
	ArticleCreateTime  string
}