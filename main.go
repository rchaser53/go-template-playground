package main

import (
	"net/http"

	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/labstack/echo"
)

type Person struct {
	UserName string
}

var tpl bytes.Buffer

func main() {
	e := echo.New()
	e.GET("/abc.js", func(c echo.Context) error {

		bs, _ := ioutil.ReadFile(`templateForGo.txt`)

		t := template.New("")
		t, _ = t.Parse(string(bs))

		t.Execute(&tpl, Person{UserName: "nyan"})

		return c.String(http.StatusOK, tpl.String())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
