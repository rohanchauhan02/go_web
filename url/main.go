package main

import (
	"fmt"
	"net/url"
)


const myurl string="https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"
func main(){
	fmt.Println("Hello GOLANG !!")
	// fmt.Println(myurl)
	res,_:=url.Parse(myurl) //parsing the url
	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.Port())
	// fmt.Println(res.Opaque)
	// fmt.Println(res.RawQuery)
	// fmt.Println(res.RawPath)
	// fmt.Println(res.Fragment)
	// fmt.Println(res.Query())
	// fmt.Println(res.RequestURI())
	qparams:=res.Query()

	fmt.Printf("the type of query params are:%T\n",qparams)
	fmt.Println(qparams["v"])
	fmt.Println(qparams["list"])
	for _,val:=range qparams{
		fmt.Println(val)
	}
	partsOfUrl:=&url.URL{ //constructing the url
		Scheme:"https",
		Host:"www.youtube.com",
		
	}
	surl:=partsOfUrl.String()
	fmt.Println(surl)
	
}