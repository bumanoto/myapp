package handle

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		data := map[string]interface{}{
			"users":          domain.FindUsers(),
			"ideas":          domain.FindIdeas(),
			"now":            time.Now().Unix(),
			"loginCompleted": sess.Values["loginCompleted"] == "completed",
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

		if password != passwordConfirm {
			data := map[string]interface{}{
				"errors": []error{errors.New("確認パスワードが一致しません")},
				"form": map[string]interface{}{
					"name":            name,
					"email":           email,
					"password":        password,
					"passwordConfirm": passwordConfirm,
				},
			}
			return c.Render(http.StatusOK, "sign_up.html", data)
		}

		existEmailUsers := domain.FindUsersByEmail(email)
		if len(existEmailUsers) > 0 {
			data := map[string]interface{}{
				"errors": []error{errors.New("登録済みのメールアドレスです")},
				"form": map[string]interface{}{
					"name":            name,
					"email":           email,
					"password":        password,
					"passwordConfirm": passwordConfirm,
				},
			}
			return c.Render(http.StatusOK, "sign_up.html", data)
		}

		cryptPassword, err := domain.PasswordEncrypt(password)
		if err != nil {
			fmt.Println("パスワード暗号化中にエラーが発生しました")
		}
		data := map[string]interface{}{}
		domain.CreateUser(name, email, cryptPassword)
		return c.Render(http.StatusOK, "sign_up_complete.html", data)
	})

	e.GET("/sign_in", func(c echo.Context) error {
		return c.Render(http.StatusOK, "sign_in.html", nil)
	})

	e.POST("/sign_in", func(c echo.Context) error {
		nameOrEmail := c.FormValue("name_or_email")
		password := c.FormValue("password")

		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["loginCompleted"] = "completed"
		sess.Save(c.Request(), c.Response())

		users := domain.FindUsersByNameOrEmail(nameOrEmail)
		if len(users) == 0 {
			data := map[string]interface{}{
				"errors": []error{errors.New("ユーザが見つかりません")},
				"form": map[string]interface{}{
					"name_or_email": nameOrEmail,
					"password":      password,
				},
			}
			return c.Render(http.StatusOK, "sign_in.html", data)
		}
		user := users[0]

		err := domain.CompareHashAndPassword(user.Password, password)
		if err != nil {
			data := map[string]interface{}{
				"errors": []error{errors.New("パスワードが一致しません")},
				"form": map[string]interface{}{
					"name_or_email": nameOrEmail,
					"password":      password,
				},
			}
			return c.Render(http.StatusOK, "sign_in.html", data)
		}

		return c.Redirect(http.StatusFound, "/")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
