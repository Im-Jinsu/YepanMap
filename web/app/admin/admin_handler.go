package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginCTRL : main
func LoginCTRL(c echo.Context) (err error) {
	// cc := c.(*cmiddleware.CustomContext)
	context := make(map[string]interface{}, 0)
	context["TITLE"] = "Yepanmap"

	return c.Render(http.StatusOK, "login.html", context)
}
