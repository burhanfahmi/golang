package main

import (
	"b47s1/connection"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	// "time"

	"github.com/labstack/echo/v4"
)

// nama dari strukturnya  adalah blog atau sama dengan nama kelas di java script
// yang membangun dari object/properties
type Blog struct {
	ID             int
	Title          string
	StartDate      string
	EndDate        string
	Description    string
	Duration       string
	IconNode       bool
	IconReact      bool
	IconJavascript bool
	IconGolang     bool
}

// menampung object yang ada pada struck Blog
// bisa disebut dummy data / menampung data sementara
// data-data yang ditampung yang kemudian data yang diisi harus sesuai dengan tipe data yang ditelah dibangun  pada strucknya
var dataBlog = []Blog{
	// {
	// 	Title:          "Dumbways Ciputat",
	// 	Description:    "Lorem Ipsum",
	// 	StartDate:      "08/06/2023",
	// 	EndDate:        "08/07/2023",
	// 	Duration:       "1 Bulan",
	// 	IconNode:       true,
	// 	IconReact:      true,
	// 	IconJavascript: true,
	// 	IconGolang:     true,
	// },
	// {
	// 	Title:          "Dumbways Depok",
	// 	Description:    "Semangat Cuy",
	// 	StartDate:      "08/06/2023",
	// 	EndDate:        "08/07/2023",
	// 	Duration:       "1 Bulan",
	// 	IconNode:       true,
	// 	IconReact:      true,
	// 	IconJavascript: false,
	// 	IconGolang:     false,
	// },
}

func main() {
	connection.DatabaseConnect()
	e := echo.New()

	// e = echo package
	// GET/POST = run the method
	// "/" = endpoint/routing (ex. localhost:5000'/' | ex. dumbways.id'/lms')
	// helloWorld = function that will run if the routes are opened
	e.Static("/public", "public")

	//routing
	// GET

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/add-project", addProject)

	//POST
	e.POST("/add-blog", addBlog)
	e.POST("/delete-project/:id", deleteBlog)
	e.POST("/edit-project/:id", editProject)
	// e.POST("/blog-delete/:id", deleteBlog)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, Startdate, enddate,duration, decription, nodejs, reactjs, javascript, golang FROM tb_blog")

	dataBlog = []Blog{}
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.StartDate, &each.EndDate, &each.Duration, &each.Description, &each.IconNode, &each.IconReact, &each.IconJavascript, &each.IconGolang)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}
		dataBlog = append(dataBlog, each)
	}
	blogs := map[string]interface{}{
		"Blogs": dataBlog,
	}
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), blogs)

}
func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{}{
	// 	"id":      id,
	// 	"Title":   "Dumbways Web App",
	// 	"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Distinctio nemo repudiandae voluptas voluptatibus modi inventore totam quaerat itaque fugiat labore! Aliquid cumque nulla iusto eaque sequi impedit rerum harum magni minus vel? Officiis quod magnam minus asperiores repellendus, autem nemo quaerat aliquid, porro nesciunt ex mollitia. Veritatis architecto voluptatem earum amet dolor enim molestias, dicta qui magni similique vero! Quis obcaecati voluptas non eum amet, mollitia, ut commodi explicabo ad praesentium debitis nemo dicta voluptatum! Voluptatum odit a voluptas, quidem temporibus inventore! Iste repellat vitae autem! Ullam expedita atque odio dolorem laudantium tempora adipisci autem nulla iste at sequi eum eaque vero blanditiis, quis tempore molestias fugiat inventore exercitationem.",
	// }

	var BlogDetail = Blog{}
	for i, data := range dataBlog {
		if id == i {
			BlogDetail = Blog{
				Title:          data.Title,
				Description:    data.Description,
				StartDate:      data.StartDate,
				EndDate:        data.EndDate,
				Duration:       data.Duration,
				IconNode:       data.IconNode,
				IconReact:      data.IconReact,
				IconJavascript: data.IconJavascript,
				IconGolang:     data.IconGolang,
			}
		}
	}
	data := map[string]interface{}{
		"Blog": BlogDetail,
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
	StartDate := c.FormValue("input-start-date")
	EndDate := c.FormValue("input-end-date")
	duration := Duration(StartDate, EndDate)
	nodeJs := c.FormValue("node")
	javaScript := c.FormValue("javascript")
	reactJs := c.FormValue("react")
	golangg := c.FormValue("golang")

	println("Tittle : " + title)
	println("Description : " + description)
	println("StartDate : " + StartDate)
	println("EndDate :" + EndDate)
	println("Duration :", duration)
	println("IconNode : " + nodeJs)
	println("IconReact : " + reactJs)
	println("IconJavascript : " + javaScript)
	println("IconGolang : " + golangg)

	//penampungan data dari struct blog
	var newBlog = Blog{
		Title:          title,
		Description:    description,
		StartDate:      time.Now().String(),
		EndDate:        time.Now().String(),
		Duration:       duration,
		IconNode:       (nodeJs == "nodejs"),
		IconJavascript: (javaScript == "javascript"),
		IconReact:      (reactJs == "reactjs"),
		IconGolang:     (golangg == "golang"),
	}

	// //append disini bertugas untuk menambahkan data newBlog kedalam slice dataBlog
	// // param 1 = dimana datany =a ditampung
	// // param 2 = data yang akan ditampung
	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/home")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/home")
}
func editProject(edit echo.Context) error {
	id, _ := strconv.Atoi(edit.Param("id"))
	fmt.Println("index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)
	return edit.Redirect(http.StatusMovedPermanently, "/add-project")
}
func Duration(StartDate, EndDate string) string {
	startTime, _ := time.Parse("2006-01-02", StartDate)
	endTime, _ := time.Parse("2006-01-02", EndDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}
