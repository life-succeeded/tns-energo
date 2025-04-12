package router

type Plugin interface {
	BasePath() string
	Register(router *Router)
}
