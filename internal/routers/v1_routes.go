package routers

import (
	"database/sql"
	"log"

	"github.com/emanuelquerty/rssagg/internal/handlers"
	"github.com/go-chi/chi"
)

type v1Router struct {
	router *chi.Mux
	Logger *log.Logger
	DBConn *sql.DB
}

func NewRouter() *v1Router {
	vr := &v1Router{}
	return vr
}

func (v1r *v1Router) Route() *chi.Mux {
	v1r.router = chi.NewRouter()
	handlerContext := handlers.HandlerContext{
		DBConn: v1r.DBConn,
		Logger: v1r.Logger,
	}

	v1r.router.Get("/healthz", handlers.HandleReadiness)
	v1r.router.Get("/err", handlers.HandleError)

	v1r.router.Post("/users", handlerContext.CreateUser)
	v1r.router.Get("/users", handlerContext.MiddlewareAuth(handlerContext.GetUser))

	v1r.router.Post("/feeds", handlerContext.MiddlewareAuth(handlerContext.CreateFeed))
	v1r.router.Get("/feeds", handlerContext.GetFeeds)

	v1r.router.Post("/feed_follows", handlerContext.MiddlewareAuth(handlerContext.CreateFeedFollows))
	v1r.router.Get("/feed_follows", handlerContext.MiddlewareAuth(handlerContext.GetFeedFollows))
	v1r.router.Delete("/feed_follows/{feedFollowID}", handlerContext.MiddlewareAuth(handlerContext.DeleteFeedFollows))

	v1r.router.Get("/posts", handlerContext.MiddlewareAuth(handlerContext.GetPosts))

	return v1r.router
}
