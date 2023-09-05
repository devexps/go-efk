package server

import (
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/http"
	"github.com/devexps/go-efk/api/gen/go/common/conf"
	v1 "github.com/devexps/go-efk/api/gen/go/helloworld/v1"
	"github.com/devexps/go-efk/helloworld/internal/service"
	"github.com/devexps/go-efk/pkg/bootstrap"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger,
	greeterSvc service.GreeterService,
) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, logging.Server(logger))
	v1.RegisterGreeterHTTPServer(srv, greeterSvc)
	return srv
}
