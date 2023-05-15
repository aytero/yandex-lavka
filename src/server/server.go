package server

import (
    "context"
    "github.com/labstack/echo/v4"
    "time"
    "yandex-team.ru/bstask/config"
    "yandex-team.ru/bstask/handler"
)

const (
    _defaultAddr            = ":8080"
    _defaultReadTimeout     = 5 * time.Second
    _defaultWriteTimeout    = 5 * time.Second
    _defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
    cfg             config.HTTP
    server          *echo.Echo
    handlers        []handler.Handler
    courierHandler  handler.Handler
    notify          chan error
    shutdownTimeout time.Duration
}

func NewServer(cfg config.HTTP, handlers ...handler.Handler) *Server {
    e := echo.New()

    s := &Server{
        server:          e,
        handlers:        handlers,
        cfg:             cfg,
        notify:          make(chan error, 1),
        shutdownTimeout: _defaultShutdownTimeout,
    }
    return s
}

func (s *Server) SetupRoutes(e *echo.Echo, handler ...handler.Handler) {
    for _, h := range handler {
        h.SetupRoutes(e)
    }
}

func (s *Server) Setup() {
    s.SetupRoutes(s.server, s.handlers...)
    s.SetupMiddleware(s.server)
}

func (s *Server) Start() {
    go func() {
        s.notify <- s.server.Start(":" + s.cfg.Port)
        close(s.notify)
    }()
}

// Notify returns server's error chan
func (s *Server) Notify() <-chan error {
    return s.notify
}

// Shutdown sets up timer for shutdown and sends a signal through context.Context
func (s *Server) Shutdown(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
    defer cancel()
    return s.server.Shutdown(ctx)
}
