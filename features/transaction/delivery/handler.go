package delivery

import (
	"fmt"
	"minibank/features/transaction/domain"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	logg "github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type trxHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := trxHandler{srv: srv}
	e.POST("/transaction", handler.AddTrx())
	e.POST("/transaction/search", handler.Filter())
	e.GET("/transaction/list", handler.GetTrx())
	e.GET("/transaction", handler.AllTrx())
}

func (th *trxHandler) AddTrx() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		if input.Credit_account == 0 || input.Debit_account == 0 || input.Amount == 0 {
			return c.JSON(http.StatusBadRequest, FailResponse("input empty"))
		}
		if input.Credit_account == input.Debit_account {
			return c.JSON(http.StatusBadRequest, FailResponse("credit and debit same"))
		}
		cnv := ToDomain(input)
		res, err := th.srv.AddTrx(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("There is problem on server."))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success to create transaction", ToResponse(res, "add")))
	}
}

func (th *trxHandler) AllTrx() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := th.srv.AllTrx()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		} else {
			return c.JSON(http.StatusOK, SuccessResponse("Success to get transaction", ToResponse(res, "all")))
		}
		return nil
	}
}

func (th *trxHandler) GetTrx() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := th.srv.Trx()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusBadRequest, FailResponse("not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
			}
		} else {
			return c.JSON(http.StatusOK, SuccessResponse("Success to list transaction", ToResponse(res, "list")))
		}
	}
}

func (th *trxHandler) Filter() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input FilterFormat
		if err := c.Bind(&input); err != nil {
			//log.Println(err.Error())
			logg.Error("error on handler", err.Error())
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		} else {
			cnv := ToDomainArray(input)
			res, err := th.srv.Search(cnv)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					logg.Error("error on handler", err.Error())
					return c.JSON(http.StatusBadRequest, FailResponse("not found"))
				} else {
					fmt.Println(input)
					return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
				}
			} else {
				fmt.Println(res)
				return c.JSON(http.StatusOK, SuccessResponse("Success to search", ToResponse(res, "all")))
			}
		}
	}
}
