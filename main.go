package main

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	pg "gopkg.in/pg.v5"
)

type Person struct {
	UserName string
}

func runServer() {
	e := echo.New()
	e.GET("/abc.js", func(c echo.Context) error {
		var tpl []byte

		response := c.Response()
		bs, _ := ioutil.ReadFile(`templateForGo.txt`)

		t := template.New("")
		t, _ = t.Parse(string(bs))

		t.Execute(response, Person{UserName: "nyan"})
		_, _ = response.Write(tpl)

		return c.String(http.StatusOK, string(tpl))
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type Hoges struct {
	Id   int64
	Name string
}

func main() {
	db := connectionDB()

	var h Hoges
	err := db.Model(&h).
		Column("hoges.*").
		Where("hoges.Id = ?", 1).
		Select()
	if err != nil {
		panic(err)
	}

	println(h.Name)

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
}

func connectionDB() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "nyan",
	})
}
