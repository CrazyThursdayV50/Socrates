package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/trace"
	"github.com/gin-gonic/gin"
)

type Server struct {
	logger        log.Logger
	tracerCreator trace.TracerCreator

	cfg *Config

	impls    Impls
	repos    Repos
	services Services
}

func New(cfg *Config, logger log.Logger, traceCreator trace.TracerCreator) *Server {
	return &Server{cfg: cfg, logger: logger, tracerCreator: traceCreator}
}

func (s *Server) Init() {
	s.initImpls()
	s.initRepos()
	s.initServices()
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	s.logger.Infof("start services on %s", addr)

	var engine = gin.Default()
	root := engine.Group("/")
	root.GET("ws", s.services.chatService.Listen)
	srv := http.Server{Handler: engine}

	wg.Add(2)
	goo.Go(func() {
		defer wg.Done()

		err := srv.Serve(l)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Errorf("start server failed: %v", err)
			return
		}
	})

	goo.Go(func() {
		defer wg.Done()

		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
		defer cancel()
		defer s.logger.Warn("SERVER EXIT ...")
		_ = srv.Shutdown(ctx)
	})
}
