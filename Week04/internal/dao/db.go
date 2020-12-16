package dao

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/pkg/errors"
	"kratos-article/internal/model"
)
//头疼啊 看了半天文档 不知道kratos 的code 怎么搞。。懵逼
//数据库文件在根目录~
var NotFound = errors.New("Not Found")

func NewDB() (db *sql.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct  paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(&cfg)
	cf = func() { db.Close() }
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.ArticleContent, err error) {
	// get data from db
	query := "SELECT * FROM article_content where article_id=?"
	art = &model.ArticleContent{}
	err = d.db.QueryRow(ctx, query, id).Scan(&art.ArticleId, &art.ArticleTitle, &art.ArticleAuthor, &art.ArticlePublishTime, &art.ArticleHits, &art.ArticleContent, &art.UserId, &art.ArticleCreateTime)
	if err != nil {
		if errors.Is(err,sql.ErrNoRows) {
			return art, errors.Wrap(NotFound, "Not Found")
		}
	}
	return
}
