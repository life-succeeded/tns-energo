package plugin

import (
	"net/http"
	"net/http/pprof"
	librouter "tns-energo/lib/http/router"
)

type PProf struct{}

func NewPProf() *PProf {
	return &PProf{}
}

func (p PProf) BasePath() string {
	return "/debug/pprof"
}

func (p PProf) Register(router *librouter.Router) {
	router.HandleGet("", librouter.WrapStdLibFunc(pprof.Index))
	router.HandleGet("/", librouter.WrapStdLibFunc(pprof.Index))
	router.HandleGet("/cmdline", librouter.WrapStdLibFunc(pprof.Cmdline))
	router.HandleGet("/profile", librouter.WrapStdLibFunc(pprof.Profile))

	router.HandleGet("/heap", loadSwHandler("heap"))
	router.HandleGet("/goroutine", loadSwHandler("goroutine"))
	router.HandleGet("/block", loadSwHandler("block"))
	router.HandleGet("/threadcreate", loadSwHandler("threadcreate"))

	router.Handle("/symbol", librouter.WrapStdLibFunc(pprof.Symbol), http.MethodGet, http.MethodPost)

	router.HandleGet("/trace", loadSwHandler("trace"))
	router.HandleGet("/mutex", loadSwHandler("mutex"))
}

func loadSwHandler(name string) librouter.Handler {
	return librouter.WrapStdLib(pprof.Handler(name))
}
