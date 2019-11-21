package main

import (
	_ "awesomeProject/zwxurl/router"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main()  {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	f, _ := os.Create("zwxurl.log")
	gin.DefaultWriter = io.MultiWriter(f)
}