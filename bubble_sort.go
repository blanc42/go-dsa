package main

import (
	"fmt"
)

func swap(a,b *int){
	temp:= *a
	*a = *b
	*b = temp
}


func main() {
	var arr = []int{12, 34,3,4,5,11,1}
	var swapped bool = true
	l := len(arr)
	for swapped {
		swapped = false
		for i:=0;i<l-1;i++ {
			if arr[i]>arr[i+1] {
				swap(&arr[i],&arr[i+1])
				swapped = true
			}	
		}
	}
	
	fmt.Print(arr)
}
