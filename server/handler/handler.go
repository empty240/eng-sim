package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する 
	    return c.String(http.StatusOK, "Hello Page")
    }
}

func JsonTest() echo.HandlerFunc {
    return func(c echo.Context) error {
	    s := []int{1,2,3}
        return c.JSON(http.StatusOK, s)
    }
}