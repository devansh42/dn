package page

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/robfig/soy"
)

//MainPage Handler
func MainPage(e echo.Context) error {
	fmt.Println(fmt.Sprint(HTMLTEMPLATES, "mainpage.htm"))
	tofu, er := soy.NewBundle().AddTemplateFile(fmt.Sprint(HTMLTEMPLATES, "mainpage.htm")).CompileToTofu()
	if er != nil {
		fmt.Print(er)
	}
	er = tofu.NewRenderer("page.main").Execute(e.Response().Writer, nil)
	if er != nil {
		fmt.Print(er)
	}
	return er
}
