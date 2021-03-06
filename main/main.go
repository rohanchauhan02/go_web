package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)
 
type Cat struct{
	Name string `json:"name"`
	Type string `json:"type`
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func getCats(c echo.Context) error{
	catName:=c.QueryParam("name")
	catType:=c.QueryParam("type")
	dataType:=c.Param("data")
 
	if dataType=="string"{
		return c.String(http.StatusOK,fmt.Sprintf("Your cat Name is:%s\nand his type is:%s\n",catName,catType))
	}
	if dataType=="json"{
		return c.JSON(http.StatusOK,map[string]string{
			"name":catName,
			"type":catType,
		})
	}
	return c.JSON(http.StatusBadRequest,map[string]string{
		"error":"You need to let us know if you want json or string data",
	})
}
func addCat(c echo.Context) error{
	cat:=Cat{}
	defer c.Request().Body.Close()
	b, err:=ioutil.ReadAll(c.Request().Body)
	if err!=nil{
		log.Printf("Failed reading the request body of addCats: %s",err)
		return c.String(http.StatusInternalServerError,"")
	}
	err=json.Unmarshal(b,&cat)
	if err!=nil{
		log.Printf("Failed unmarshaling in addCats:%s ",err)
		return c.String(http.StatusInternalServerError,"")
	}
	log.Printf("This is your cat: %v",cat)
	return c.String(http.StatusOK,"we got your cat")
}
func main() {
  
	fmt.Println("Welcome to the server")
	// Echo instance
	e := echo.New()

	// Routes

	//Read
	e.GET("/", hello)
	e.GET("/cats/:data",getCats)
	//Write

	e.POST("/cats",addCat)
	// Start server
	e.Start(":1323")
}
