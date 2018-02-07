package main

import (
	"net/http"

	"io/ioutil"
	"text/template"

	"github.com/labstack/echo"
)

type Person struct {
	UserName string
}

func main() {
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
