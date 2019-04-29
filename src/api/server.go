package api

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func Serve() {
	InitConfig()

	http.HandleFunc("/simple/.health_check", SimpleHealthCheck)
	http.HandleFunc("/simple/.health_check/deep", SimpleDeepHealthCheck)

	// See https://dev.to/davidsbond/golang-debugging-memory-leaks-using-pprof-5di8
	//http.HandleFunc("/debug/pprof/", pprof.Index)
	//http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//
	//http.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	//http.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	//http.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	//http.Handle("/debug/pprof/block", pprof.Handler("block"))

	log.Printf("start http server by port '%s'", config.Http.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Http.Port), nil); err != nil {
		log.Fatalf("failed to start http server by error '%#v'", err)
	}
}
