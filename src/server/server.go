package server

import (
    "context"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/gommon/log"
    "time"
)

const (
    _defaultReadTimeout     = 5 * time.Second
    _defaultWriteTimeout    = 5 * time.Second
    _defaultAddr            = ":8080"
    _defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
    server          *echo.Echo
    handlers        []Handler
    courierHandler  Handler
    notify          chan error
    shutdownTimeout time.Duration
}

func NewServer(handlers ...Handler) *Server {

    e := echo.New()

    // todo config loglevel
    e.Logger.SetLevel(log.DEBUG)
    //e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.Logger())
    //e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    //    AllowOrigins: []string{"*"},
    //    AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
    //    AllowMethods: []string{echo.GET, echo.POST},
    //}))
    //e.Validator = NewValidator()

    //rate := limiter.Rate{
    //	Limit:  10,
    //	Period: time.Second,
    //}
    //e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

    //for _, h := range handlers {
    //    routes.SetupRoutes(e, h)
    //}
    // todo logging by middleware
    //e.Logger.Debug()
    s := &Server{
        server:          e,
        handlers:        handlers,
        notify:          make(chan error, 1),
        shutdownTimeout: _defaultShutdownTimeout,
    }

    return s

}

func (s *Server) Start() {
    go func() {
        s.notify <- s.server.Start(":8080")
        close(s.notify)
    }()
}

// Notify returns server's error chan
func (s *Server) Notify() <-chan error {
    return s.notify
}

// Shutdown sets up timer for shutdown and sends a signal through context.Context
func (s *Server) Shutdown(ctx context.Context) error {
    //ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
    //defer cancel()

    return s.server.Shutdown(ctx)
}

/*
type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
*/
