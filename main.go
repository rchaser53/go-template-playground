package main

import (
	"net/http"

	"html/template"
	"os"

	"github.com/labstack/echo"
)

type Person struct {
	UserName string
}

func main() {
	e := echo.New()
	e.GET("/abc.js", func(c echo.Context) error {

		file, err := os.Open(`template.js`)
		if err != nil {
			// Openエラー処理
		}
		defer file.Close()

		t := template.New("template.js")
		t, _ = t.Parse("hello {{.UserName}}!")
		p := Person{UserName: "Astaxie"}
		t.Execute(os.Stdout, p)

		output := "testmessage"
		file.Write(([]byte)(output))

		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
