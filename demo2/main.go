package main

/*
int number_add(int a, int b) {
    return a+b;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.number_add(11, 10))
}
