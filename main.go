package main

import "fmt"

func main() {
	fmt.Println("Hello World,1992")
	var array1 [3]*string

	array2 := [3]*string{new(string), new(string), new(string)}

	*array2[0] = "RED"
	*array2[1] = "GREEN"
	*array2[2] = "BLUE"

	array1 = array2

	for i := 0; i < 3; i++ {
		fmt.Printf("array1[%d] = %s\n", i, *array1[i])
	}

}
