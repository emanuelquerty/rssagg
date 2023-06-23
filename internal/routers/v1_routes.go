package routers

import (
	"github.com/emanuelquerty/rssagg/internal/handlers"
	"github.com/go-chi/chi"
)

type RouterContext struct {
	router     *chi.Mux
	handlerCtx handlers.HandlerContext
}

func NewRouterContext(handlerCtx handlers.HandlerContext) *RouterContext {
	routerCtx := &RouterContext{
		router:     chi.NewRouter(),
		handlerCtx: handlerCtx,
	}
	return routerCtx
}

func (rctx *RouterContext) Route() *chi.Mux {
	hctx := rctx.handlerCtx

	rctx.router.Get("/healthz", handlers.HandleReadiness)
	rctx.router.Get("/err", handlers.HandleError)

	rctx.router.Post("/users", hctx.CreateUser)
	rctx.router.Get("/users", hctx.MiddlewareAuth(hctx.GetUser))

	rctx.router.Post("/feeds", hctx.MiddlewareAuth(hctx.CreateFeed))
	rctx.router.Get("/feeds", hctx.GetFeeds)

	rctx.router.Post("/feed_follows", hctx.MiddlewareAuth(hctx.CreateFeedFollows))
	rctx.router.Get("/feed_follows", hctx.MiddlewareAuth(hctx.GetFeedFollows))
	rctx.router.Delete("/feed_follows/{feedFollowID}", hctx.MiddlewareAuth(hctx.DeleteFeedFollows))

	rctx.router.Get("/posts", hctx.MiddlewareAuth(hctx.GetPosts))

	return rctx.router
}
