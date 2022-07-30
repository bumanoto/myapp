package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var Db *gorm.DB

type User struct {
	gorm.Model
	Name string
}

func init() {
	initDb()
}

func initDb() {
	dialector := mysql.Open("myappuser:myapppass@tcp(127.0.0.1:3306)/myapp_database?charset=utf8&parseTime=true")
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
	fmt.Println("db connected!!")
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		var users []User
		Db.Find(&users)
		//data := "Hello"

		data := map[string]interface{}{"users": users}

		return c.Render(http.StatusOK, "index.html", data)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
