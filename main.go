package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e = echo package
	// GET/POST = run the method
	// "/" = endpoint/routing (ex. localhost:5000'/' | ex. dumbways.id'/lms')
	// helloWorld = function that will run if the routes are opened
	e.Static("/public", "public")

	//routing
	e.GET("/hello", helloWorld)
	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/add-project", addProject)
	e.POST("/add-blog", addBlog)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func blog(c echo.Context) error {
	data := map[string]interface{}{
		"Login": true,
	}

	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}
func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":      id,
		"Title":   "Dumbways Web App",
		"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Distinctio nemo repudiandae voluptas voluptatibus modi inventore totam quaerat itaque fugiat labore! Aliquid cumque nulla iusto eaque sequi impedit rerum harum magni minus vel? Officiis quod magnam minus asperiores repellendus, autem nemo quaerat aliquid, porro nesciunt ex mollitia. Veritatis architecto voluptatem earum amet dolor enim molestias, dicta qui magni similique vero! Quis obcaecati voluptas non eum amet, mollitia, ut commodi explicabo ad praesentium debitis nemo dicta voluptatum! Voluptatum odit a voluptas, quidem temporibus inventore! Iste repellat vitae autem! Ullam expedita atque odio dolorem laudantium tempora adipisci autem nulla iste at sequi eum eaque vero blanditiis, quis tempore molestias fugiat inventore exercitationem.",
	}
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}
func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func addBlog(c echo.Context) error {
	title := c.FormValue("input-title")
	description := c.FormValue("input-textarea")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")

	println("Tittle : " + title)
	println("Description : " + description)
	println("Start date : " + startDate)
	println("End date :" + endDate)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
