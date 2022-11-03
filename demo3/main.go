package main

//#include "number.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.number_add(11, 10))
}
