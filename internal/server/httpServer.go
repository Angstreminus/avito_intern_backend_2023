package server

import (
	"database/sql"
	"log"

	"github.com/Angstreminus/avito_intern_backend_2023/config"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/controller"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/repository"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Angstreminus/avito_intern_backend_2023/docs"
)

type HttpServer struct {
	config *config.Config
	router *gin.Engine
}

func NewHttpServer(config *config.Config, dbhandler *sql.DB) *HttpServer {

	segmentRepository := repository.NewSegmentRepository(dbhandler)
	segmentsUserRep := repository.NewSegmentUserRepository(dbhandler)
	segmentService := service.NewSegmentService(segmentRepository)
	segmentsUserService := service.NewSegmentUserService(segmentsUserRep)
	segmentController := controller.NewSegmentController(segmentService)
	segmetsUserController := controller.NewUserSegmentController(segmentsUserService)
	router := gin.Default()

	router.POST("/segments", segmentController.CreateSegment)
	router.DELETE("/segments/:id", segmentController.DeleteSegment)
	router.POST("/segments/users", segmetsUserController.CreateSegmentsUsers)
	router.PUT("/segments/users/:id", segmetsUserController.EditUserSegment)
	router.GET("/segments/users:id", segmetsUserController.GetSegmentsNamesByUserID)
	router.GET("/swagger/doc", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return &HttpServer{
		config: config,
		router: router,
	}
}

func (hS HttpServer) Start() {
	err := hS.router.Run(":8080")
	if err != nil {
		log.Fatalf("Error while start up: %v", err)
	}
}
