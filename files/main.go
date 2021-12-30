package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Printf("Welcome to files in golang")
	content:="This needs to go in a file- Learn Go Fast!!"
	file,err:=os.Create("./myFile.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	errorHandler(err)
	length, err:=io.WriteString(file,content)
	errorHandler(err)
	fmt.Println("Length is: ",length)
	defer file.Close()
	readFile("./myFile.txt")
}
func readFile(filename string){
	databyte,err:=ioutil.ReadFile(filename)
	errorHandler(err)
	fmt.Println(databyte)
	fmt.Println("the file content goes from here:", string(databyte))
}
func errorHandler(err error){
	if err!=nil{
		panic(err)
	}
}