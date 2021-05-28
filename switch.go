//switch

package main

import (
	"fmt"
)
func main(){
	fmt.Println("Enter a key:")
	var key int
	fmt.Scanln(&key)
	switch(key){
	case 1:
		fmt.Println("\nONE")
	case 2:
		fmt.Println("\nTWO")
	case 3:
		fmt.Println("\nTHREE")
	case 4:
	  	fmt.Println("\nFOUR")
	case 5:
		fmt.Println("\nFIVE")
	case 6:
		fmt.Println("\nSIX")
	default:
		fmt.Println("\nnot in the option!!1")
		}
	}
