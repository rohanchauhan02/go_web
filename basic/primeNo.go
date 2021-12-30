package main

import (
	"fmt"
	"math"
)

func printPrimeNumbers(num1, num2 int){
	if num1<2 || num2<2 {
		fmt.Println("number must be greater than 2")
		return
	}

	for num1<=num2{
		isPrime:=true //shorthand syntex
		for i:=2;i<=int(math.Sqrt(float64(num1)));i++{
			if num1%i==0{
				isPrime=false
				break
			}
		}
		if isPrime{
			fmt.Println(num1)
		}
		num1++
	}
}


func main(){
	printPrimeNumbers(5,15)
}
