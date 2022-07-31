package handle

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"myapp/domain"
	"net/http"
	"time"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Handle() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/static/", "./static/")

	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"users": domain.FindUsers(),
			"ideas": domain.FindIdeas(),
			"now":   time.Now().Unix(),
		}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.GET("/sign_up", func(c echo.Context) error {
		data := map[string]interface{}{}
		return c.Render(http.StatusOK, "sign_up.html", data)
	})

	e.POST("/sign_up", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		passwordConfirm := c.FormValue("password_confirm")
		domain.CreateUser(name, email, password, passwordConfirm)
		data := map[string]interface{}{}
		return c.Render(http.StatusOK, "sign_up.html", data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
