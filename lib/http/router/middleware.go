package router

type Middleware func(c Context, h Handler) Handler
