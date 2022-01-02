package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"io"
	"os"

)
type User struct{
	Name string `json;"name" xml:"name" form:"name" query:"name`
	Email string `json:"email" xml:"email" form:"email" query:="email"`
}

func hello(c echo.Context) error{
	return c.String(http.StatusOK,"Hello World!")
}
func getUser(c echo.Context) error{
	//User ID from path users/:id
	id:=c.Param("id")
	return c.String(http.StatusOK,id)
}
func show(c echo.Context)error{
	//Get team and member from the query string
	team:=c.QueryParam("team")
	member:=c.QueryParam("member")
	return c.String(http.StatusOK,"team:"+team+", member:"+member)
}
func save(c echo.Context) error{
	//Get name and email
	name:=c.FormValue("name")
	email:=c.FormValue("email")
	return c.String(http.StatusOK,"name:" + name + ", email:" + email)
}

func info(c echo.Context) error{
	//Get name
	name:=c.FormValue("name")
	//Get avatar
	avatar,err:=c.FormFile("avatar")
	if err!=nil{
		return err
	}

	//Source

	src,err:=avatar.Open()
	if err!=nil{
		return err
	}
	defer src.Close()

	//Destination

	dst ,err:=os.Create(avatar.Filename)
	if err!=nil{
		return err
	}
	defer dst.Close()

	//Copy

	if _,err=io.Copy(dst,src);err!=nil{
		return err
	}
	return c.HTML(http.StatusOK,"<b>Thank you!"+name+"<b>")

}
func createUser(c echo.Context) error{
	u:=new(User)
	if err:=c.Bind(u);err!=nil{
		return err
	}
	return c.JSON(http.StatusCreated,u)
	// return c.XML(http.StatusCreated, u)
}
func main(){
	fmt.Println("Server is running on port 1323")
	e:=echo.New()

	//route
	e.GET("/",hello)
	e.GET("/users/:id",getUser)
	e.GET("/show",show)
	e.POST("/save",save)
	e.POST("/info",info)
	e.POST("/users",createUser)
	// e.post("/users",saveUser)
	// e.PUT("/users/:id",updateUser)
	// e.DELETE("/users/:id",deleteUser)

//start server
	e.Logger.Fatal(e.Start(":8000"))

}