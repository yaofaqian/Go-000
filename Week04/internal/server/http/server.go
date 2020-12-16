package http

import (
	"net/http"

	pb "kratos-article/api"
	"kratos-article/internal/model"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svc pb.ArticleServer

// New new a bm server.
func New(s pb.ArticleServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterArticleBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/kratos-article")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	articleContent := &model.ArticleContent{
		ArticleId:          1,
		ArticleTitle:       "苹果M1芯片相较于华为麒麟9000芯片是否被严重高估了？",
		ArticleAuthor:      "互联网砖瓦匠",
		ArticlePublishTime: "2020-12-15",
		ArticleHits:        110,
		ArticleContent:     "虽然还没有拿到M1实机，但只要对比华为最新的手机芯片——麒麟9000 就清楚了，麒麟9000作为一款拥有153",
		UserId:             1,
		ArticleCreateTime:  "2020-12-15",
	}
	c.JSON(articleContent, nil)
}
