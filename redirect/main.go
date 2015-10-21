package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Fprintln(os.Stdout,"std output")
	fmt.Fprintln(os.Stderr,"std err output")
}