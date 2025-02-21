package main

<<<<<<< HEAD
import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World, 1992")

=======
import "fmt"

func main() {
	fmt.Println("Hello World,1992")
>>>>>>> 787eb1239df8a7dca640dec8fbd25867d1f9728c
	var array1 [3]*string

	array2 := [3]*string{new(string), new(string), new(string)}

	*array2[0] = "RED"
	*array2[1] = "GREEN"
	*array2[2] = "BLUE"

	array1 = array2

	for i := 0; i < 3; i++ {
		fmt.Printf("array1[%d] = %s\n", i, *array1[i])
	}

<<<<<<< HEAD
	// 添加HTTP处理函数
	http.HandleFunc("/", helloHandler)

	// 启动HTTP服务器
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你好, this is the root endpoint!")
=======
>>>>>>> 787eb1239df8a7dca640dec8fbd25867d1f9728c
}
