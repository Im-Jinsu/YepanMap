package route

import (
	"io"
	"text/template"

	"github.com/im-jinsu/yepanmap/web/app/admin"
	"github.com/im-jinsu/yepanmap/web/app/nintendo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Route set path
func Route() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.HTTPSWWWRedirect())

	// Set templates
	e.Renderer = SetTemplates()

	// Set static
	e.Static("/static", "static")

	// Old version
	e.GET("/", nintendo.OldMainPageCTRL)
	e.GET("/nintendo", nintendo.OldMainPageCTRL)
	e.GET("/old_map", nintendo.OldMainPageCTRL)

	// Main
	// ------------ Controller
	e.GET("/map", nintendo.MainPageCTRL)
	// ------------ AJAX
	// e.GET("/nintendo.prc", nintendo.MainPageCTRL)

	// Admin
	adminRouter := e.Group("/shop")
	// Login
	// ------------ Controller
	adminRouter.GET("/login", admin.LoginCTRL)

	return e
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// SetTemplates set templates
func SetTemplates() *TemplateRenderer {
	// Set templates
	templates := []string{
		// Old_version
		"templates/old_main.html",
		"templates/old_base.html",
		// Main
		"templates/base.html",
		"templates/main.html",
		"templates/admin/login.html",
	}
	t, _ := template.New("").ParseFiles(templates...)
	renderer := &TemplateRenderer{
		templates: t,
	}
	return renderer
}
