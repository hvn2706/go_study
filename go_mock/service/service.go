package service

import (
	"context"
	"gomock/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type AlbumInfo struct {
	Title  string
	Artist string
	Price  float64
}

func RunServer() {
	router := gin.Default()

	router.POST("/albums", PostAlbum)
	router.GET("/albums/:id", ReturnAlbum)
	router.PUT("/albums/edit", EditAlbum)
	router.DELETE("/albums/:id", DeleteAlbum)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.FatalLogger.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	quit_signal := <-quit

	logger.Log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.FatalLogger.Fatalf("Server forced to shutdown: %s\n", err)
	}

	logger.InfoLogger.Printf("Server exiting: %v\n", quit_signal)
}
