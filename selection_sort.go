package main

import (
	"fmt"
)

func swap(a,b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var arr = []int{14,12,3,5,7,2} 
	fmt.Println(arr)
	l := len(arr)
	for i:= 0;i<l;i++ {
		for j:=i+1;j<l;j++ {
			if(arr[i]>arr[j]){
				swap(&arr[i],&arr[j])
			}
		}
	}	
	fmt.Println(arr)
}
