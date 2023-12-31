package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(v,w,x,y,z int ) int {
	return v + w + x + y + z
}
