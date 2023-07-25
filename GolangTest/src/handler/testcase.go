package handler

import (
	"fmt"
	"golangtest/src/payload/request"
	"golangtest/src/payload/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCountrySpend(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.useCase.GetCountrySpend(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "ok",
		Message: "success",
		Data:    result,
	})
}
func (h *Handler) GetTotalCardType(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.useCase.GetCardTypeTotal(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "ok",
		Message: "success",
		Data:    result,
	})
}
func (h *Handler) GetDataByCountry(c echo.Context) error {
	ctx := c.Request().Context()
	country := c.Param("country")
	if country == "" || country == ":country" {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: "Invalid Request",
			Data:    "Country not found!",
		})
	}
	fmt.Println(country)
	result, err := h.useCase.GetDataByCountry(ctx, country)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "ok",
		Message: "success",
		Data:    result,
	})
}
func (h *Handler) CreateDataUsers(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.ReqUsers{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: "Invalid Request",
			Data:    err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: "Invalid Request",
			Data:    err.Error(),
		})
	}
	result, err := h.useCase.CreateDataUser(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.GlobalResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.GlobalResponse{
		Status:  "ok",
		Message: "success",
		Data:    result,
	})
}
