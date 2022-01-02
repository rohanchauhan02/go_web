package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Hello golang!!")
	// PerformGetRequest()
}
func PerformPostJsonRequest(){
	const myurl=""
}

func PerformGetRequest(){
	const myurl="http://aviabird.com"
	res,err:=http.Get(myurl)
	if err!=nil{
		 panic(err)
	} 
	defer res.Body.Close()
	fmt.Println("Status code :",res.Status)

	fmt.Println("content length is:",res.ContentLength)
	var responseString strings.Builder
	content,_:=ioutil.ReadAll(res.Body) //byte format
	byteCount,_:=responseString.Write(content)
	fmt.Println("ByteCount is:",byteCount) 
	// fmt.Println(string(content))   
	fmt.Println(responseString.String() ) 
} 