package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gofiber-idp/server/api"
	"gofiber-idp/server/config"
	"gofiber-idp/server/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
)

type Server struct {
	app  *fiber.App
	db   *gorm.DB
	port int
}

func NewServer(port int, logLevel logger.LogLevel) *Server {
	logger.InitLogger(logLevel)

	return &Server{
		app:  fiber.New(),
		port: port,
	}
}

func (s *Server) Start() error {
	s.registerDB()
	s.registerMiddleware()
	s.registerApi()

	err := s.app.Listen(fmt.Sprintf(":%d", s.port))
	return err
}

func (s *Server) registerDB() {
	c, err := config.LoadConfig()

	if err != nil {
		logger.LogError("")
	}

	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{
		Logger: logger2.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger2.Config{
				LogLevel: logger.Info, // Log level
				Colorful: true,        // Disable color
			},
		),
	})
	if err != nil {
		panic(err)
	}

	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)

	s.db = db
}

func (s *Server) registerApi() {
	api.InitUserApi(s.app, s.db)
}

func (s *Server) registerMiddleware() {
	s.app.Use(RecoverAndStackTraceMiddleware)
	s.app.Use(RequestPrintMiddleware)
	s.app.Use(ResponsePrintMiddleware)
}
