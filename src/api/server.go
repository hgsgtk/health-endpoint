package api

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	InitConfig()

	http.HandleFunc("/simple/.health_check", SimpleHealthCheck)
	http.HandleFunc("/simple/.health_check/deep", SimpleDeepHealthCheck)

	log.Printf("start http server by port '%s'", config.Http.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Http.Port), nil); err != nil {
		log.Fatalf("failed to start http server by error '%#v'", err)
	}
}
