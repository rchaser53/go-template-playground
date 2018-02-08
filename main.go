package main

import (
	"fmt"
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

func main() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "papillon_development",
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "select count(Id) from users")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
