package main

import (
	"fmt"
)

func FindPerfectnum (x int) {
	count1:=0
	for i:=1;i<x;i++{
		if x%i==0{
			count1+=i
		}
	}
	if count1==x{
		fmt.Println(x)
	}
}
func FindPrimernum(x int)  {
	flag:=1
	for i:=2;i*i<=x;i++{
		if x%i==0{
			flag=0
			break
		}
	}
	if flag==1 {
		fmt.Println(x)
	}
}
func main()  {
	fmt.Println("完美数有:")
	for i2:=2;i2<=123456;i2++{
		go FindPerfectnum(i2)
	}
	fmt.Println("素数有")
	for i3:=1;i3<=123456;i3++{
		 FindPrimernum(i3)
	}
}
