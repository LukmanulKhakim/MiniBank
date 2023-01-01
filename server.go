package main

import (
	"minibank/config"
	dAccount "minibank/features/account/delivery"
	rAccount "minibank/features/account/repository"
	sAccount "minibank/features/account/services"
	dTrx "minibank/features/transaction/delivery"
	rTrx "minibank/features/transaction/repository"
	sTrx "minibank/features/transaction/services"
	"minibank/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	mdlAct := rAccount.New(db)
	serAct := sAccount.New(mdlAct)
	dAccount.New(e, serAct)

	mdlTrx := rTrx.New(db)
	serTrx := sTrx.New(mdlTrx)
	dTrx.New(e, serTrx)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8000"))

}
