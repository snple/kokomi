package web

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snple/beacon/core"
	"go.uber.org/zap"
)

type WebService struct {
	cs *core.CoreService

	auth     *AuthService
	user     *UserService
	device   *DeviceService
	slot     *SlotService
	source   *SourceService
	tag      *TagService
	constant *ConstService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts webOptions
}

func NewWebService(cs *core.CoreService, opts ...WebApiOption) (*WebService, error) {
	ctx, cancel := context.WithCancel(cs.Context())

	s := &WebService{
		cs:     cs,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultWebOptions(),
	}

	for _, opt := range extrawebOptions {
		opt.apply(&s.dopts)
	}

	for _, opt := range opts {
		opt.apply(&s.dopts)
	}

	auth, err := newAuthService(s)
	if err != nil {
		return nil, err
	}
	s.auth = auth

	s.user = newUserService(s)
	s.device = newDeviceService(s)
	s.slot = newSlotService(s)
	s.source = newSourceService(s)
	s.tag = newTagService(s)
	s.constant = newConstService(s)

	return s, nil
}

func (s *WebService) GetAuth() *AuthService {
	return s.auth
}

func (s *WebService) GetUser() *UserService {
	return s.user
}

func (s *WebService) Register(router gin.IRouter) {
	s.auth.register(router)
	s.user.register(router)
	s.device.register(router)
	s.slot.register(router)
	s.source.register(router)
	s.tag.register(router)
	s.constant.register(router)
}

func (s *WebService) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()
}

func (s *WebService) Stop() {
	s.cancel()
	s.closeWG.Wait()

	s.Logger().Sync()
}

func (s *WebService) Core() *core.CoreService {
	return s.cs
}

func (s *WebService) Context() context.Context {
	return s.ctx
}

func (s *WebService) Logger() *zap.Logger {
	return s.dopts.logger
}

type webOptions struct {
	debug         bool
	logger        *zap.Logger
	jwtSecretKey  string
	jwtTimeout    time.Duration
	jwtMaxRefresh time.Duration
}

func defaultWebOptions() webOptions {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment(): %v", err)
	}

	return webOptions{
		debug:         false,
		logger:        logger,
		jwtTimeout:    time.Hour * 24,
		jwtMaxRefresh: time.Hour * 24,
	}
}

type WebApiOption interface {
	apply(*webOptions)
}

var extrawebOptions []WebApiOption

type funcWebApiOption struct {
	f func(*webOptions)
}

func (fdo *funcWebApiOption) apply(do *webOptions) {
	fdo.f(do)
}

func newFuncWebApiOption(f func(*webOptions)) *funcWebApiOption {
	return &funcWebApiOption{
		f: f,
	}
}

func WithDebug(debug bool) WebApiOption {
	return newFuncWebApiOption(func(o *webOptions) {
		o.debug = debug
	})
}

func WithLogger(logger *zap.Logger) WebApiOption {
	return newFuncWebApiOption(func(o *webOptions) {
		o.logger = logger
	})
}

func WithJwtSecretKey(key string) WebApiOption {
	return newFuncWebApiOption(func(o *webOptions) {
		o.jwtSecretKey = key
	})
}
