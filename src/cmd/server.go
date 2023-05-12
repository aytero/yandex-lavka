package main

import (
    "github.com/labstack/echo/v4"
    "log"
    "yandex-team.ru/bstask/config"
    "yandex-team.ru/bstask/database"
    "yandex-team.ru/bstask/handler"
    "yandex-team.ru/bstask/repository"
    "yandex-team.ru/bstask/routes"
    "yandex-team.ru/bstask/usecase"
)

func main() {

    conf, err := config.NewConfig("config/.config.env")
    if err != nil {
        log.Fatalln("Initial configuration failed")
    }

    db, err := database.New(conf.URL)
    if err != nil {
        log.Fatalln("failed db conn")
    }

    //d := db.New()
    //db.AutoMigrate(d)

    // todo courier repo n usecase
    orderRepo := repository.NewOrderRepository(db)
    orderUsecase := usecase.NewOrderUsecase(orderRepo)

    // handler
    //e := server.New()
    //e.SetupServer(handlers)

    //if err := goose.Run(command, db, *dir, arguments...); err != nil {
    //	log.Fatalf("goose %v: %v", command, err)
    //}

    //e := setupServer()
    e := setupServer(orderUsecase)
    e.Logger.Fatal(e.Start(":8080"))
    // todo graceful shutdown
}

/*
func setupServer() *echo.Echo {
	e := echo.New()
	routes.SetupRoutes(e)
	return e
}
*/

func setupServer(ordUsecase *usecase.OrderUsecase) *echo.Echo {
    e := echo.New()

    //todo: handle the error!
    //c, _ := handlers.NewContainer()

    // Middleware
    //e.Use(middleware.Logger())
    //e.Use(middleware.Recover())

    // Start server
    orderHandler := handler.NewOrderHandler(e, ordUsecase)
    routes.SetupRoutes(e, orderHandler)
    log.Println("after setup")

    return e
}
