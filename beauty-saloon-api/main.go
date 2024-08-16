package main

import (
	"log"
	"net/http"
	"beauty-saloon-api/controllers"
	"beauty-saloon-api/database"
	"beauty-saloon-api/repository"
	"beauty-saloon-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a conexão com o banco de dados
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}
	defer db.Close()

	// Inicializa os repositórios
	appointmentRepo := repository.NewAppointmentRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	saloonRepo := repository.NewSaloonRepository(db)

	// Inicializa os serviços
	appointmentService := services.NewAppointmentService(appointmentRepo)
	customerService := services.NewCustomerService(customerRepo)
	saloonService := services.NewSaloonService(saloonRepo)

	// Inicializa os controladores
	appointmentController := controllers.NewAppointmentController(appointmentService)
	customerController := controllers.NewCustomerController(customerService)
	saloonController := controllers.NewSaloonController(saloonService)

	// Cria o roteador Gin
	router := gin.Default()

	// Define as rotas
	api := router.Group("/api")
	{
		appointments := api.Group("/appointments")
		{
			appointments.POST("/", appointmentController.CreateAppointment)
		}

		customers := api.Group("/customers")
		{
			customers.POST("/", customerController.CreateCustomer)
		}

		saloons := api.Group("/saloons")
		{
			saloons.POST("/", saloonController.CreateSaloon)
		}
	}

	// Inicia o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
