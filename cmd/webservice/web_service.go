package webservice

import (
	"log"
	"net/http"


	"github.com/Erwin011895/shorty-challenge/cmd/webservice/public/handler"
	"github.com/Erwin011895/shorty-challenge/cmd/webservice/public/router"
	"github.com/Erwin011895/shorty-challenge/internal/config"
	"github.com/Erwin011895/shorty-challenge/internal/module"
	"github.com/Erwin011895/shorty-challenge/internal/component"
)

var (
	srv *http.Server
)

// Start initiate the webservice binary
func Start() (err error) {
	s := component.InitSharedComponent()
	m := module.InitModules(s)
	h := handler.NewHandler(s, m)

	router := router.Init(h, s)

	srv = &http.Server{
		Addr:    ":" + config.Get().HttpPort,
		Handler: router,
	}
	
	log.Println("start webservice. listening on", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln("failed starting web on", srv.Addr, err)
	}

	return err
}
