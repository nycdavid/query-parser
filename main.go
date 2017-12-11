package main

import (
	"fmt"
	"net/url"

	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
}

func (ctx *CustomContext) Qvalues() url.Values {
	return ctx.Request().URL.Query()
}

func QueryParser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cc := &CustomContext{ctx}
		return next(cc)
	}
}

func main() {
	e := echo.New()
	e.Use(QueryParser)
	e.GET("/", Home)
	e.Logger.Fatal(e.Start(":1323"))
	fmt.Println("Listening on port 1323...")
}

func Home(ctx echo.Context) error {
	cc := ctx.(*CustomContext)
	fmt.Println(cc.Qvalues())
	return nil
}
