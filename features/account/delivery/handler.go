package delivery

import (
	"errors"
	"fmt"
	"minibank/features/account/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type accountHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := accountHandler{srv: srv}
	e.POST("/account", handler.AddAccount())
	e.GET("/account/:id", handler.GetMyAccount())
	e.GET("/account/list", handler.GetAllAccount())

}

func (ah *accountHandler) AddAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		if strings.TrimSpace(input.Name) == "" {
			return c.JSON(http.StatusBadRequest, FailResponse("name empty"))
		}
		if input.Balance == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("balance empty"))
		}
		cnv := ToDomain(input)
		res, err := ah.srv.AddAccount(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "already") {
				c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse("There is problem on server."))
			}
		} else {

			return c.JSON(http.StatusCreated, SuccessResponse("Success to create account", ToResponse(res, "add")))
		}
		return nil
	}

}

func (ah *accountHandler) GetAllAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input bool
		details := c.QueryParam("details")
		if strings.Contains(details, "true") {
			input = true
		} else if strings.Contains(details, "false") {
			input = false
		}
		res, err := ah.srv.AllAccount(input)
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		} else {
			if input == true {
				fmt.Println(input)
				fmt.Println(res)
				return c.JSON(http.StatusOK, SuccessResponse("Success to list account", ToResponse(res, "all")))
			} else if input == false {
				fmt.Println(input)
				fmt.Println(res)
				return c.JSON(http.StatusOK, SuccessResponse("Success to list account", ToResponse(res, "no")))
			}
		}

		return nil
	}
}

func (ah *accountHandler) GetMyAccount() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errors.New("cannot convert id")
		}
		res, err := ah.srv.MyAccount(uint(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusBadRequest, FailResponse("not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
			}
		} else {
			return c.JSON(http.StatusOK, SuccessResponse("Success show my content", ToResponse(res, "add")))
		}
	}
}
