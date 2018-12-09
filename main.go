package main

import (
	"bookapi/pkg/logging"
	"bookapi/routers"
	"fmt"
	"net/http"

	"bookapi/pkg/setting"
)

func main() {
	logging.Info("test")
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
