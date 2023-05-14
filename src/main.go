package main

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "os"
    "yandex-team.ru/bstask/config"
    "yandex-team.ru/bstask/courier"
    "yandex-team.ru/bstask/database"
    "yandex-team.ru/bstask/handler"
    "yandex-team.ru/bstask/order"
    "yandex-team.ru/bstask/repository"
    "yandex-team.ru/bstask/routes"
    "yandex-team.ru/bstask/usecase"
)

func main() {

    conf, err := config.NewConfig("config/.config.env")
    if err != nil {
        log.Error().Msg("Initial configuration failed")
        return
    }

    db, err := database.New(conf.URL)
    if err != nil {
        log.Error().Msg(fmt.Sprintf("failed db conn: %v", err))
        return
    }

    //d := db.New()
    //db.AutoMigrate(d)

    // todo APP setup

    // todo courier repo n usecase
    orderRepo := repository.NewOrderRepository(db)
    orderUsecase := usecase.NewOrderUsecase(orderRepo)

    courierRepo := repository.NewCourierRepository(db)
    courierUsecase := usecase.NewCourierUsecase(courierRepo)

    // handler
    //e := server.New()
    //e.SetupServer(handlers)

    //if err := goose.Run(command, db, *dir, arguments...); err != nil {
    //	log.Fatalf("goose %v: %v", command, err)
    //}

    // todo important
    e := setupServer(orderUsecase, courierUsecase)
    e.Logger.Fatal(e.Start(":8080"))
    //routes.SetupRoutes(e, orderHandler)
    // todo graceful shutdown

    /*
       // todo experimental
       orderHandler := handler.NewOrderHandler(orderUsecase)
       echoServ := server.NewServer(orderHandler) // , courierHandler
       //e.SetupServer(handlers)
       echoServ.Start()

       // Waiting signal
       interrupt := make(chan os.Signal, 1)
       signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

       select {
       case sig := <-interrupt:
           log.Printf("shutting down with signal: %s\n", sig)
       case err = <-echoServ.Notify():
           log.Printf("server err: %s", err)
       }

       ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
       defer func() {
           cancel()
           //defer db.Close()
       }()
       //    defer server.Stop()
       //    defer sc.Close()
       //}()

       // Shutdown
       err = echoServ.Shutdown(ctx)
       if err != nil {
           log.Printf("server shutdown err: %s", err)
       }
    */
    /*
    	srv := new(server.Server)
    	go func() {
    		if err := srv.Run(viper.GetString("port"), app(db)); err != nil {
    			logrus.Fatalf("error occured while running http server: %s", err.Error())
    		}
    	}()

    	logrus.Print("Server Started")

    	quit := make(chan os.Signal, 1)
    	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
    	<-quit

    	logrus.Print("Server Shutting Down")

    	if err := srv.Shutdown(context.Background()); err != nil {
    		logrus.Errorf("error occured on server shutting down: %s", err.Error())
    	}

    	if err := db.Close(); err != nil {
    		logrus.Errorf("error occured on db connection close: %s", err.Error())
    	}

    */
}

func setupServer(ordUsecase order.Usecase, curUsecase courier.Usecase) *echo.Echo {
    e := echo.New()

    //c, err := handlers.NewContainer()

    // Middleware
    e.Use(middleware.Recover())

    // move loggger to main
    logger := zerolog.New(os.Stdout)
    zerolog.SetGlobalLevel(zerolog.DebugLevel)

    e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
        LogURI:    true,
        LogStatus: true,
        LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
            logger.Info().
                Str("URI", v.URI).
                Int("status", v.Status).
                Msg("request")

            return nil
        },
    }))

    // Start server
    orderHandler := handler.NewOrderHandler(ordUsecase)
    courierHandler := handler.NewCourierHandler(curUsecase)
    routes.SetupRoutes(e, orderHandler)
    routes.SetupRoutes(e, courierHandler)
    log.Debug().Msg("after setup")

    return e
}
