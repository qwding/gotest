package main

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "time"
)

// Handler
func hello(c echo.Context) error {
	// panic("hello panic")
  time.Sleep(time.Second*2)
  return c.String(http.StatusOK, "Hello, World!")
}

func main() {
  // Echo instance
  e := echo.New()
  e.Logger.SetPrefix("fdafdsf")


  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Use(useTime)

  e.Use(order)

  // Routes
  e.GET("/", hello)
  e.Logger.Error("this is error")

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}


func order(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    println("order")
    err := next(c)
    println("order end")
    return err
  }
}


func useTime(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    start := time.Now()
    println("use Time")
    err := next(c)

    end := time.Now()

    println("use:", end.Sub(start).Seconds())

    return err
  }
}