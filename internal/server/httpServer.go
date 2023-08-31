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
	config                 *config.Config
	router                 *gin.Engine
	segmentController      *controller.SegmentController
	userSegmentsController *controller.UserSegmentController
}

func InitHttpServer(config *config.Config, dbhandler *sql.DB) *HttpServer {

	segmentRepository := repository.NewSegmentRepository(dbhandler)
	userSegmentsRep := repository.NewUserSegmentRepository(dbhandler)
	segmentService := service.NewSegmentService(segmentRepository)
	userSegmentsService := service.NewUserSegmentService(userSegmentsRep)
	segmentController := controller.NewSegmentController(segmentService)
	userSegmetsController := controller.NewUserSegmentController(userSegmentsService)
	router := gin.Default()

	router.POST("/segments", segmentController.CreateSegment)
	router.DELETE("/segments/:id", segmentController.DeleteSegment)
	router.POST("/user_segments", userSegmetsController.CreateUserSegment)
	router.PUT("/user_segments/:id", userSegmetsController.EditUserSegment)
	router.GET("/user_segments/:id", userSegmetsController.GetSegmentsNamesByUserID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return &HttpServer{
		config:                 config,
		router:                 router,
		segmentController:      segmentController,
		userSegmentsController: userSegmetsController,
	}
}

func (hS HttpServer) Start() {
	err := hS.router.Run(":8080")
	if err != nil {
		log.Fatalf("Error while start up: %v", err)
	}
}
