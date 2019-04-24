package output

import "fmt"

func PrintRow(row []int){
	
	for _, v := range row {
		fmt.Printf("%d\t", v)		
	}
	fmt.Printf("\n")		
}
