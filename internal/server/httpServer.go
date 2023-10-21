package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	Config *config.Config
	Router *gin.Engine
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
		Config: config,
		Router: router,
	}
}

func (hS *HttpServer) Start() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: hS.Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while starting server: %v\n", err)
		}
	}()

	log.Println("Server is on 8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Signal to end")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error to shutdown: %v\n", err)
	}

	log.Println("Server disabled in time")
}
