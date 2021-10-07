package endpoint

import (
	"github.com/labstack/echo/v4"
	"github.com/onurozerdal/hepsiburada-case/api/service"
	"net/http"
)

type ApiController struct {
	service service.ApiService
}

func NewApiController(apiService service.ApiService) *ApiController {
	return &ApiController{service: apiService}
}

func (controller *ApiController) BrowsingHistories(ctx echo.Context) error {
	userId := ctx.QueryParams().Get("user-id")
	r, err := controller.service.BrowsingHistories(userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, r)
}

func (controller *ApiController) BestsellerProducts(ctx echo.Context) error {
	userId := ctx.QueryParams().Get("user-id")
	r, err := controller.service.BestsellerProducts(userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, r)
}

func (controller *ApiController) DeleteHistory(ctx echo.Context) error {
	userId := ctx.QueryParams().Get("user-id")
	productId := ctx.QueryParams().Get("product-id")
	history, err := controller.service.DeleteHistory(userId, productId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, history)
}
