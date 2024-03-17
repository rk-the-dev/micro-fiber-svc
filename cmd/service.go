package cmd

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	mLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rk-the-dev/micro-fiber-svc/api"
	"github.com/rk-the-dev/micro-fiber-svc/app"
	"github.com/sirupsen/logrus"
)

type Microservice struct {
	name   string
	log    *logrus.Logger
	port   int
	server *fiber.App
}

func NewSVC(name string, port int) *Microservice {
	return &Microservice{name: name, port: port}
}
func (m *Microservice) Start() {
	app.Logger.Infof("Starting %s the service", m.name)
	m.InitApp()
	api.RegisterAPI(m.server)
	logRegisteredRoutes(m.server)
	m.server.Listen(":" + strconv.Itoa(m.port))
}
func (m *Microservice) InitApp() {
	m.server = fiber.New()
	m.server.Use(mLogger.New())
	m.server.Get("/ping", api.Ping)
	m.server.Get("/health", api.Health)
}
func logRegisteredRoutes(app *fiber.App) {
	fmt.Println("Registered Routes:")
	for _, route := range app.Stack() {
		for _, r := range route {
			fmt.Printf("Method: %s, Path: %s\n", r.Method, r.Path)
		}
	}
}
