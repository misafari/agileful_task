package main

import (
	"agileful_task/internal/core/service/query_svc"
	"agileful_task/internal/handler/http_handler"
	"agileful_task/internal/handler/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=admin password=admin dbname=sampleDb port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
		PrepareStmt:            true,
		CreateBatchSize:        100,
	})

	if err != nil {

	}

	queriesDatabaseRepository := repository.NewQueriesDatabaseRepository(db, true)
	appQueriesService := query_svc.NewQueriesService(queriesDatabaseRepository)
	handler := http_handler.NewQueriesHttpHandler(appQueriesService)

	app.Get("/queries/count", handler.GetQueriesCountHandler)

	// ex: http://127.0.0.1:3000/queries?sort=desc&type=SELECT&pagination.page=2&pagination.page_size=10
	app.Get("/queries", handler.GetQueriesHandler)

	log.Fatal(app.Listen(":3000"))
}
