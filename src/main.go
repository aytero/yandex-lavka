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
	courierHandler "yandex-team.ru/bstask/courier/handler"
	courierRepo "yandex-team.ru/bstask/courier/repository"
	courierUsecase "yandex-team.ru/bstask/courier/usecase"
	"yandex-team.ru/bstask/database"
	orderHandler "yandex-team.ru/bstask/order/handler"
	orderRepo "yandex-team.ru/bstask/order/repository"
	orderUsecase "yandex-team.ru/bstask/order/usecase"
	"yandex-team.ru/bstask/server"
)

func main() {
	// todo add app pkg: app.Run()
	// todo mv config to cmd, add default configs
	// todo add panic recovery etc
	// todo tests

	// todo remove sleep ?
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

	ordRepo := orderRepo.NewOrderRepository(db)
	ordUsecase := orderUsecase.NewOrderUsecase(ordRepo)

	courRepo := courierRepo.NewCourierRepository(db)
	courUsecase := courierUsecase.NewCourierUsecase(courRepo)

	ordHandler := orderHandler.NewOrderHandler(ordUsecase)
	courHandler := courierHandler.NewCourierHandler(courUsecase)

	echoServ := server.NewServer(cfg.HTTP, ordHandler, courHandler)
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
