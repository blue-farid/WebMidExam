package main

import (
	c "github.com/blue-farid/WebMidExam/config"
	h "github.com/blue-farid/WebMidExam/handler"
	r "github.com/blue-farid/WebMidExam/repository"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	_, err := c.ReadConfig()
	e := echo.New()
	if err == nil {
		r.InitDB()
		h.InitRoutes(e)
	}
}
