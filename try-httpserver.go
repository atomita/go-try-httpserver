package tryHttpserver

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

func Start(port int64) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/helloworld", func(c echo.Context) error {
		data := map[string]interface{}{"result": "hellow, world", "code": 123}
		encoded, _ := json.Marshal(data)
		return c.String(http.StatusOK, string(encoded))
	})

	p := strconv.FormatInt(port, 10)

	fmt.Print("serve 127.0.0.1:" + p)
	e.Run(standard.New(":" + p))
}
