package http

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	engine *gin.Engine

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts httpServerOptions
}

func NewHttpServer(ctx context.Context, opts ...HttpServerOption) (*HttpServer, error) {
	ctx, cancel := context.WithCancel(ctx)

	engine := gin.New()

	engine.RedirectFixedPath = true

	{
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowAllOrigins = true
		corsConfig.AddAllowHeaders("Authorization")
		corsConfig.AddAllowMethods("OPTIONS")

		engine.Use(cors.New(corsConfig))
	}

	s := &HttpServer{
		engine: engine,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultHttpServerOptions(),
	}

	for _, opt := range extraHttpServerOptions {
		opt.apply(&s.dopts)
	}

	for _, opt := range opts {
		opt.apply(&s.dopts)
	}

	s.engine.Use(gin.RecoveryWithWriter(os.Stderr))

	if s.dopts.debug {
		s.engine.Use(gin.LoggerWithWriter(os.Stdout))
	}

	return s, nil
}

func (s *HttpServer) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	server := &http.Server{
		Addr:      s.dopts.addr,
		Handler:   s.engine.Handler(),
		TLSConfig: s.dopts.tlsConfig.Clone(),
	}

	go func() {
		if s.dopts.certFile != "" && s.dopts.keyFile != "" {
			s.Logger().Sugar().Infof("%s: listening and serving HTTPS on %s", s.dopts.appName, s.dopts.addr)

			if err := server.ListenAndServeTLS(s.dopts.certFile, s.dopts.keyFile); err != nil {
				if err != http.ErrServerClosed {
					s.Logger().Sugar().Fatalf("listen: %v", err)
				}
			}
		} else {
			s.Logger().Sugar().Infof("%s: listening and serving HTTP on %s", s.dopts.appName, s.dopts.addr)

			if err := server.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					s.Logger().Sugar().Fatalf("listen: %v", err)
				}
			}
		}
	}()

	<-s.ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		s.Logger().Sugar().Errorf("server shutdown: %v", err)
	}
}

func (s *HttpServer) Stop() {
	s.cancel()
	s.closeWG.Wait()

	s.Logger().Sync()
}

func (s *HttpServer) Engine() *gin.Engine {
	return s.engine
}

func (s *HttpServer) Context() context.Context {
	return s.ctx
}

func (s *HttpServer) Logger() *zap.Logger {
	return s.dopts.logger
}

type httpServerOptions struct {
	appName           string
	debug             bool
	logger            *zap.Logger
	addr              string
	certFile, keyFile string
	tlsConfig         *tls.Config
}

func defaultHttpServerOptions() httpServerOptions {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment(): %v", err)
	}

	return httpServerOptions{
		debug:  false,
		logger: logger,
		addr:   ":8006",
	}
}

type HttpServerOption interface {
	apply(*httpServerOptions)
}

var extraHttpServerOptions []HttpServerOption

type funcHttpServerOption struct {
	f func(*httpServerOptions)
}

func (fdo *funcHttpServerOption) apply(do *httpServerOptions) {
	fdo.f(do)
}

func newFuncHttpServerOption(f func(*httpServerOptions)) *funcHttpServerOption {
	return &funcHttpServerOption{
		f: f,
	}
}

func WithAppName(name string) HttpServerOption {
	return newFuncHttpServerOption(func(o *httpServerOptions) {
		o.appName = name
	})
}

func WithDebug(debug bool) HttpServerOption {
	return newFuncHttpServerOption(func(o *httpServerOptions) {
		o.debug = debug
	})
}

func WithLogger(logger *zap.Logger) HttpServerOption {
	return newFuncHttpServerOption(func(o *httpServerOptions) {
		o.logger = logger
	})
}

func WithAddr(addr string) HttpServerOption {
	return newFuncHttpServerOption(func(o *httpServerOptions) {
		o.addr = addr
	})
}

func WithTLS(certFile, keyFile string) HttpServerOption {
	return newFuncHttpServerOption(func(o *httpServerOptions) {
		o.certFile = certFile
		o.keyFile = keyFile
	})
}

func WithTLSConfig(tlsConfig *tls.Config) HttpServerOption {
	return newFuncHttpServerOption(func(o *httpServerOptions) {
		o.tlsConfig = tlsConfig
	})
}
