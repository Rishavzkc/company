package router

import (
	"company/config"
	"company/controller"
	"company/database"
	"company/repository"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CompanyRouting() {

	configs := config.NewConfig()
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.Database,
	)

	db, err := database.SetupDatabaseConnection(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := database.CloseDatabaseConnection(db); err != nil {
			log.Fatal(err)
		}
	}()

	r := gin.Default()

	dbController := repository.New(db)
	companyController := controller.NewCompanyController(dbController)

	companyRouter := r.Group("/company")
	{
		companyRouter.POST("/", companyController.CreateCompany)
		companyRouter.GET("/", companyController.GetAllCompanies)
		companyRouter.GET("/:id", companyController.GetCompany)
		companyRouter.PUT("/", companyController.UpdateCompany)
		companyRouter.DELETE("/:id", companyController.DeleteCompany)
	}

	if err := r.Run(configs.ServiceHost); err != nil {
		log.Fatal("failure at running server: %w", err)
	}
}
