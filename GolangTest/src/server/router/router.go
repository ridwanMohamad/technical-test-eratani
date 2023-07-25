package router

import (
	"golangtest/src/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitializeRouter(server *echo.Echo, handler *handler.Handler) {
	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "service up")
	})
	v1 := server.Group("/api/v1")
	{
		v1.GET("/get-country-spend", handler.GetCountrySpend)
		v1.GET("/get-total-cc-by-type", handler.GetTotalCardType)
		v1.GET("/:country/get-data", handler.GetDataByCountry)
		v1.POST("/create-users", handler.CreateDataUsers)
	}
}
