package main

import (
	"fmt"
	"humanInfoCollection/pkg"
	"humanInfoCollection/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", pkg.HTTPPort),
		Handler:        router,
		ReadTimeout:    pkg.ReadTimeout,
		WriteTimeout:   pkg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
