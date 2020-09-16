package nintendo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// OldMainPageCTRL : main
func OldMainPageCTRL(c echo.Context) (err error) {
	context := make(map[string]interface{}, 0)
	context["TITLE"] = "YepanMap - Find Nintendo Switch"
	context["HEADER"] = "Nintendo"

	return c.Render(http.StatusOK, "old_main.html", context)
}

// MainPageCTRL : main
func MainPageCTRL(c echo.Context) (err error) {
	// cc := c.(*cmiddleware.CustomContext)
	context := make(map[string]interface{}, 0)
	context["TITLE"] = "YepanMap - Find Nintendo Switch"
	context["HEADER"] = "Nintendo"

	return c.Render(http.StatusOK, "main.html", context)
}
