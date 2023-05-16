package main

import (
    "context"
    "fmt"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "os"
    "os/signal"
    "syscall"
    "time"
    "yandex-team.ru/bstask/config"
    "yandex-team.ru/bstask/database"
    "yandex-team.ru/bstask/handler"
    "yandex-team.ru/bstask/repository"
    "yandex-team.ru/bstask/server"
    "yandex-team.ru/bstask/usecase"
)

func main() {
    time.Sleep(5 * time.Second)
    zerolog.SetGlobalLevel(zerolog.ErrorLevel)

    cfg, err := config.NewConfig("config/.config.env")
    if err != nil {
        log.Error().Msg("Initial configuration failed")
        cfg, err = config.NewConfig("default")
        if err != nil {
            log.Fatal().Msg("Default configuration failed")
        }
    }

    db, err := database.New(cfg.URL)
    if err != nil {
        log.Error().Msg(fmt.Sprintf("failed db conn: %v", err))
        return
    }

    orderRepo := repository.NewOrderRepository(db)
    orderUsecase := usecase.NewOrderUsecase(orderRepo)

    courierRepo := repository.NewCourierRepository(db)
    courierUsecase := usecase.NewCourierUsecase(courierRepo)

    orderHandler := handler.NewOrderHandler(orderUsecase)
    courierHandler := handler.NewCourierHandler(courierUsecase)

    echoServ := server.NewServer(cfg.HTTP, orderHandler, courierHandler)
    echoServ.Setup()

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
        defer db.Close()
        //    defer server.Stop()
    }()

    // Shutdown
    err = echoServ.Shutdown(ctx)
    if err != nil {
        log.Error().Msg(fmt.Sprintf("server shutdown err: %v", err))
    }
}
