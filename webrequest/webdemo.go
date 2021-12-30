package main

import (
	"fmt"
	"net/http"
	// "ioutil"
)
const url="https://aviabird.com"
func main(){
	fmt.Println("Hello World!")
	res,err:=http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Response is of type :%T\n",res) //pointer
	defer res.Body.Close()
	databyte,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		panic(err)
	}
	content:=string(databyte)
	fmt.Printf(content)
}