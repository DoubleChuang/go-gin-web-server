package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.LoadHTMLGlob("resources/*.html")
	router.Static("/static", "resources/static")
	router.GET("/", GetIndex)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	fmt.Printf("Port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
